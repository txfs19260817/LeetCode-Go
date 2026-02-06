package designbankingsystem

import (
	"container/heap"
	"fmt"
	"sort"
)

const cashbackDelay = 86400000

type balanceEntry struct {
	timestamp int
	balance   int
}

type account struct {
	id        string
	balance   int
	outgoing  int
	createdAt int
	mergedAt  int
	history   []balanceEntry
	payments  map[string]struct{}
}

type payment struct {
	id        string
	accountID string
	cashback  int
	due       int
	received  bool
}

type cashbackItem struct {
	due       int
	paymentID string
}

type cashbackHeap []cashbackItem

func (h cashbackHeap) Len() int           { return len(h) }
func (h cashbackHeap) Less(i, j int) bool { return h[i].due < h[j].due }
func (h cashbackHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *cashbackHeap) Push(x interface{}) {
	*h = append(*h, x.(cashbackItem))
}

func (h *cashbackHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type BankingSystem struct {
	accounts   map[string]*account
	payments   map[string]*payment
	cashbacks  cashbackHeap
	paymentSeq int
}

func NewBankingSystem() *BankingSystem {
	return &BankingSystem{
		accounts: make(map[string]*account),
		payments: make(map[string]*payment),
	}
}

func (bs *BankingSystem) CreateAccount(timestamp int, accountID string) bool {
	bs.processCashbacks(timestamp)
	if _, exists := bs.accounts[accountID]; exists {
		return false
	}
	acct := &account{
		id:        accountID,
		createdAt: timestamp,
		mergedAt:  -1,
		history:   []balanceEntry{{timestamp: timestamp, balance: 0}},
		payments:  make(map[string]struct{}),
	}
	bs.accounts[accountID] = acct
	return true
}

func (bs *BankingSystem) Deposit(timestamp int, accountID string, amount int) int {
	bs.processCashbacks(timestamp)
	acct := bs.getActiveAccount(accountID)
	if acct == nil {
		return -1
	}
	bs.setBalance(acct, timestamp, acct.balance+amount)
	return acct.balance
}

func (bs *BankingSystem) Transfer(timestamp int, sourceAccountID string, targetAccountID string, amount int) int {
	bs.processCashbacks(timestamp)
	if sourceAccountID == targetAccountID {
		return -1
	}
	source := bs.getActiveAccount(sourceAccountID)
	if source == nil {
		return -1
	}
	target := bs.getActiveAccount(targetAccountID)
	if target == nil {
		return -1
	}
	if source.balance < amount {
		return -1
	}

	bs.setBalance(source, timestamp, source.balance-amount)
	bs.setBalance(target, timestamp, target.balance+amount)
	source.outgoing += amount
	return source.balance
}

func (bs *BankingSystem) TopSpenders(timestamp int, n int) []string {
	bs.processCashbacks(timestamp)
	type entry struct {
		id       string
		outgoing int
	}
	entries := make([]entry, 0, len(bs.accounts))
	for _, acct := range bs.accounts {
		if acct.mergedAt != -1 {
			continue
		}
		entries = append(entries, entry{id: acct.id, outgoing: acct.outgoing})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].outgoing != entries[j].outgoing {
			return entries[i].outgoing > entries[j].outgoing
		}
		return entries[i].id < entries[j].id
	})

	if n > len(entries) {
		n = len(entries)
	}
	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, fmt.Sprintf("%s(%d)", entries[i].id, entries[i].outgoing))
	}
	return result
}

func (bs *BankingSystem) Pay(timestamp int, accountID string, amount int) string {
	bs.processCashbacks(timestamp)
	acct := bs.getActiveAccount(accountID)
	if acct == nil {
		return ""
	}
	if acct.balance < amount {
		return ""
	}
	bs.setBalance(acct, timestamp, acct.balance-amount)
	acct.outgoing += amount

	bs.paymentSeq++
	paymentID := fmt.Sprintf("payment%d", bs.paymentSeq)
	cashback := amount * 2 / 100
	due := timestamp + cashbackDelay
	p := &payment{
		id:        paymentID,
		accountID: accountID,
		cashback:  cashback,
		due:       due,
	}
	bs.payments[paymentID] = p
	acct.payments[paymentID] = struct{}{}
	heap.Push(&bs.cashbacks, cashbackItem{due: due, paymentID: paymentID})
	return paymentID
}

func (bs *BankingSystem) GetPaymentStatus(timestamp int, accountID string, paymentID string) string {
	bs.processCashbacks(timestamp)
	acct := bs.getActiveAccount(accountID)
	if acct == nil {
		return ""
	}
	p, exists := bs.payments[paymentID]
	if !exists || p.accountID != accountID {
		return ""
	}
	if p.received {
		return "CASHBACK_RECEIVED"
	}
	return "IN_PROGRESS"
}

func (bs *BankingSystem) MergeAccounts(timestamp int, accountID1 string, accountID2 string) bool {
	bs.processCashbacks(timestamp)
	if accountID1 == accountID2 {
		return false
	}
	acct1 := bs.getActiveAccount(accountID1)
	if acct1 == nil {
		return false
	}
	acct2 := bs.getActiveAccount(accountID2)
	if acct2 == nil {
		return false
	}

	bs.setBalance(acct1, timestamp, acct1.balance+acct2.balance)
	acct1.outgoing += acct2.outgoing
	for paymentID := range acct2.payments {
		p := bs.payments[paymentID]
		if p != nil {
			p.accountID = acct1.id
			acct1.payments[paymentID] = struct{}{}
		}
	}
	acct2.payments = nil
	acct2.mergedAt = timestamp
	return true
}

func (bs *BankingSystem) GetBalance(timestamp int, accountID string, timeAt int) int {
	bs.processCashbacks(timestamp)
	acct, exists := bs.accounts[accountID]
	if !exists {
		return -1
	}
	if timeAt < acct.createdAt {
		return -1
	}
	if acct.mergedAt != -1 && timeAt >= acct.mergedAt {
		return -1
	}
	history := acct.history
	idx := sort.Search(len(history), func(i int) bool {
		return history[i].timestamp > timeAt
	})
	if idx == 0 {
		return -1
	}
	return history[idx-1].balance
}

func (bs *BankingSystem) getActiveAccount(accountID string) *account {
	acct, exists := bs.accounts[accountID]
	if !exists || acct.mergedAt != -1 {
		return nil
	}
	return acct
}

func (bs *BankingSystem) setBalance(acct *account, timestamp int, balance int) {
	acct.balance = balance
	n := len(acct.history)
	if n > 0 && acct.history[n-1].timestamp == timestamp {
		acct.history[n-1].balance = balance
		return
	}
	acct.history = append(acct.history, balanceEntry{timestamp: timestamp, balance: balance})
}

func (bs *BankingSystem) processCashbacks(timestamp int) {
	for bs.cashbacks.Len() > 0 {
		item := bs.cashbacks[0]
		if item.due > timestamp {
			return
		}
		heap.Pop(&bs.cashbacks)
		p := bs.payments[item.paymentID]
		if p == nil || p.received {
			continue
		}
		acct := bs.accounts[p.accountID]
		if acct == nil || acct.mergedAt != -1 {
			continue
		}
		bs.setBalance(acct, item.due, acct.balance+p.cashback)
		p.received = true
	}
}
