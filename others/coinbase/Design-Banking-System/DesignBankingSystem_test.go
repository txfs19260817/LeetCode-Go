package designbankingsystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankingSystemExample1(t *testing.T) {
	bs := NewBankingSystem()

	assert.True(t, bs.CreateAccount(1, "account1"))
	assert.False(t, bs.CreateAccount(2, "account1"))
	assert.True(t, bs.CreateAccount(3, "account2"))
	assert.Equal(t, -1, bs.Deposit(4, "non-existing", 2700))
	assert.Equal(t, 2700, bs.Deposit(5, "account1", 2700))
	assert.Equal(t, -1, bs.Transfer(6, "account1", "account2", 2701))
	assert.Equal(t, 2500, bs.Transfer(7, "account1", "account2", 200))
}

func TestBankingSystemExample2(t *testing.T) {
	bs := NewBankingSystem()

	assert.True(t, bs.CreateAccount(1, "acc1"))
	assert.True(t, bs.CreateAccount(2, "acc2"))
	assert.True(t, bs.CreateAccount(3, "acc3"))
	assert.Equal(t, 1000, bs.Deposit(4, "acc1", 1000))
	assert.Equal(t, 500, bs.Deposit(5, "acc2", 500))
	assert.Equal(t, 0, bs.Deposit(6, "acc3", 0))
	assert.Equal(t, -1, bs.Transfer(7, "acc1", "acc1", 100))
	assert.Equal(t, 700, bs.Transfer(8, "acc1", "acc2", 300))
	assert.Equal(t, 0, bs.Transfer(9, "acc2", "acc3", 800))
	assert.Equal(t, 0, bs.Transfer(10, "acc1", "acc3", 700))
}

func TestBankingSystemExample3(t *testing.T) {
	bs := NewBankingSystem()

	assert.True(t, bs.CreateAccount(1, ""))
	assert.False(t, bs.CreateAccount(2, ""))
	assert.True(t, bs.CreateAccount(3, "user@123"))
	assert.True(t, bs.CreateAccount(4, "user-account_1"))
	assert.Equal(t, 1000000, bs.Deposit(5, "", 1000000))
	assert.Equal(t, 2500, bs.Deposit(6, "user@123", 2500))
	assert.Equal(t, -1, bs.Transfer(7, "", "nonexistent", 500))
	assert.Equal(t, -1, bs.Transfer(8, "nonexistent", "user@123", 100))
	assert.Equal(t, 750000, bs.Transfer(9, "", "user@123", 250000))
	assert.Equal(t, 0, bs.Transfer(10, "", "user-account_1", 750000))
	assert.Equal(t, 50, bs.Deposit(11, "", 50))
}

func TestBankingSystemTopSpenders(t *testing.T) {
	bs := NewBankingSystem()

	assert.True(t, bs.CreateAccount(1, "account3"))
	assert.True(t, bs.CreateAccount(2, "account2"))
	assert.True(t, bs.CreateAccount(3, "account1"))
	assert.Equal(t, 2000, bs.Deposit(4, "account2", 2000))
	assert.Equal(t, 3000, bs.Deposit(5, "account3", 3000))
	assert.Equal(t, 4000, bs.Deposit(6, "account1", 4000))
	assert.Equal(t, []string{"account1(0)", "account2(0)", "account3(0)"}, bs.TopSpenders(7, 3))
	assert.Equal(t, 3500, bs.Transfer(8, "account3", "account2", 500))
	assert.Equal(t, 2500, bs.Transfer(9, "account3", "account1", 1000))
	assert.Equal(t, 500, bs.Transfer(10, "account1", "account2", 2500))
	assert.Equal(t, []string{"account1(2500)", "account3(1500)", "account2(0)"}, bs.TopSpenders(11, 3))
}

func TestBankingSystemPayments(t *testing.T) {
	bs := NewBankingSystem()

	assert.True(t, bs.CreateAccount(1, "account1"))
	assert.True(t, bs.CreateAccount(2, "account2"))
	assert.Equal(t, 2000, bs.Deposit(3, "account1", 2000))
	assert.Equal(t, "payment1", bs.Pay(4, "account1", 1000))
	assert.Equal(t, "payment2", bs.Pay(100, "account1", 1000))
	assert.Equal(t, "", bs.GetPaymentStatus(101, "non-existing", "payment1"))
	assert.Equal(t, "", bs.GetPaymentStatus(102, "account2", "payment1"))
	assert.Equal(t, "IN_PROGRESS", bs.GetPaymentStatus(103, "account1", "payment1"))
	assert.Equal(t, []string{"account1(2000)", "account2(0)"}, bs.TopSpenders(104, 2))
	assert.Equal(t, 100, bs.Deposit(86400003, "account1", 100))
	assert.Equal(t, "CASHBACK_RECEIVED", bs.GetPaymentStatus(86400004, "account1", "payment1"))
	assert.Equal(t, 220, bs.Deposit(86400005, "account1", 100))
	assert.Equal(t, 320, bs.Deposit(86400099, "account1", 100))
	assert.Equal(t, 440, bs.Deposit(86400100, "account1", 100))
}

func TestBankingSystemMergesAndBalances(t *testing.T) {
	bs := NewBankingSystem()

	assert.True(t, bs.CreateAccount(1, "account1"))
	assert.True(t, bs.CreateAccount(2, "account2"))
	assert.Equal(t, 2000, bs.Deposit(3, "account1", 2000))
	assert.Equal(t, 2000, bs.Deposit(4, "account2", 2000))
	assert.Equal(t, "payment1", bs.Pay(5, "account2", 300))
	assert.Equal(t, 1500, bs.Transfer(6, "account1", "account2", 500))
	assert.False(t, bs.MergeAccounts(7, "account1", "non-existing"))
	assert.False(t, bs.MergeAccounts(8, "account1", "account1"))
	assert.True(t, bs.MergeAccounts(9, "account1", "account2"))
	assert.Equal(t, 3800, bs.Deposit(10, "account1", 100))
	assert.Equal(t, -1, bs.Deposit(11, "account2", 100))
	assert.Equal(t, "", bs.GetPaymentStatus(12, "account2", "payment1"))
	assert.Equal(t, "IN_PROGRESS", bs.GetPaymentStatus(13, "account1", "payment1"))
	assert.Equal(t, -1, bs.GetBalance(14, "account2", 1))
	assert.Equal(t, -1, bs.GetBalance(15, "account2", 9))
	assert.Equal(t, 3800, bs.GetBalance(16, "account1", 11))
	assert.Equal(t, 3906, bs.Deposit(86400005, "account1", 100))
}
