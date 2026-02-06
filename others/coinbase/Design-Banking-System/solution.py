from typing import Dict, List, Optional, Set, Tuple
import bisect
import heapq


CASHBACK_DELAY = 86400000


class _Account:
    def __init__(self, account_id: str, created_at: int) -> None:
        self.id = account_id
        self.balance = 0
        self.outgoing = 0
        self.created_at = created_at
        self.merged_at: Optional[int] = None
        self.history: List[Tuple[int, int]] = [(created_at, 0)]
        self.payments: Set[str] = set()

    def set_balance(self, timestamp: int, balance: int) -> None:
        self.balance = balance
        if self.history and self.history[-1][0] == timestamp:
            self.history[-1] = (timestamp, balance)
        else:
            self.history.append((timestamp, balance))


class _Payment:
    def __init__(
        self, payment_id: str, account_id: str, cashback: int, due: int
    ) -> None:
        self.id = payment_id
        self.account_id = account_id
        self.cashback = cashback
        self.due = due
        self.received = False


class BankingSystem:
    def __init__(self) -> None:
        self._accounts: Dict[str, _Account] = {}
        self._payments: Dict[str, _Payment] = {}
        self._cashbacks: List[Tuple[int, str]] = []
        self._payment_seq = 0

    def createAccount(self, timestamp: int, accountId: str) -> bool:
        self._process_cashbacks(timestamp)
        if accountId in self._accounts:
            return False
        self._accounts[accountId] = _Account(accountId, timestamp)
        return True

    def deposit(self, timestamp: int, accountId: str, amount: int) -> int:
        self._process_cashbacks(timestamp)
        account = self._get_active(accountId)
        if account is None:
            return -1
        account.set_balance(timestamp, account.balance + amount)
        return account.balance

    def transfer(
        self,
        timestamp: int,
        sourceAccountId: str,
        targetAccountId: str,
        amount: int,
    ) -> int:
        self._process_cashbacks(timestamp)
        if sourceAccountId == targetAccountId:
            return -1
        source = self._get_active(sourceAccountId)
        target = self._get_active(targetAccountId)
        if source is None or target is None:
            return -1
        if source.balance < amount:
            return -1

        source.set_balance(timestamp, source.balance - amount)
        target.set_balance(timestamp, target.balance + amount)
        source.outgoing += amount
        return source.balance

    def topSpenders(self, timestamp: int, n: int) -> List[str]:
        self._process_cashbacks(timestamp)
        entries = [
            (account.outgoing, account.id)
            for account in self._accounts.values()
            if account.merged_at is None
        ]
        entries.sort(key=lambda item: (-item[0], item[1]))
        if n > len(entries):
            n = len(entries)
        return [f"{account_id}({outgoing})" for outgoing, account_id in entries[:n]]

    def pay(self, timestamp: int, accountId: str, amount: int) -> str:
        self._process_cashbacks(timestamp)
        account = self._get_active(accountId)
        if account is None or account.balance < amount:
            return ""
        account.set_balance(timestamp, account.balance - amount)
        account.outgoing += amount

        self._payment_seq += 1
        payment_id = f"payment{self._payment_seq}"
        cashback = amount * 2 // 100
        due = timestamp + CASHBACK_DELAY
        payment = _Payment(payment_id, accountId, cashback, due)
        self._payments[payment_id] = payment
        account.payments.add(payment_id)
        heapq.heappush(self._cashbacks, (due, payment_id))
        return payment_id

    def getPaymentStatus(self, timestamp: int, accountId: str, paymentId: str) -> str:
        self._process_cashbacks(timestamp)
        account = self._get_active(accountId)
        if account is None:
            return ""
        payment = self._payments.get(paymentId)
        if payment is None or payment.account_id != accountId:
            return ""
        return "CASHBACK_RECEIVED" if payment.received else "IN_PROGRESS"

    def mergeAccounts(self, timestamp: int, accountId1: str, accountId2: str) -> bool:
        self._process_cashbacks(timestamp)
        if accountId1 == accountId2:
            return False
        account1 = self._get_active(accountId1)
        account2 = self._get_active(accountId2)
        if account1 is None or account2 is None:
            return False

        account1.set_balance(timestamp, account1.balance + account2.balance)
        account1.outgoing += account2.outgoing
        for payment_id in account2.payments:
            payment = self._payments.get(payment_id)
            if payment is not None:
                payment.account_id = account1.id
                account1.payments.add(payment_id)
        account2.payments.clear()
        account2.merged_at = timestamp
        return True

    def getBalance(self, timestamp: int, accountId: str, timeAt: int) -> int:
        self._process_cashbacks(timestamp)
        account = self._accounts.get(accountId)
        if account is None:
            return -1
        if timeAt < account.created_at:
            return -1
        if account.merged_at is not None and timeAt >= account.merged_at:
            return -1
        timestamps = [entry[0] for entry in account.history]
        idx = bisect.bisect_right(timestamps, timeAt)
        if idx == 0:
            return -1
        return account.history[idx - 1][1]

    def _get_active(self, account_id: str) -> Optional[_Account]:
        account = self._accounts.get(account_id)
        if account is None or account.merged_at is not None:
            return None
        return account

    def _process_cashbacks(self, timestamp: int) -> None:
        while self._cashbacks and self._cashbacks[0][0] <= timestamp:
            due, payment_id = heapq.heappop(self._cashbacks)
            payment = self._payments.get(payment_id)
            if payment is None or payment.received:
                continue
            account = self._accounts.get(payment.account_id)
            if account is None or account.merged_at is not None:
                continue
            account.set_balance(due, account.balance + payment.cashback)
            payment.received = True


if __name__ == "__main__":
    bs = BankingSystem()
    assert bs.createAccount(1, "account1") is True
    assert bs.createAccount(2, "account1") is False
    assert bs.createAccount(3, "account2") is True
    assert bs.deposit(4, "non-existing", 2700) == -1
    assert bs.deposit(5, "account1", 2700) == 2700
    assert bs.transfer(6, "account1", "account2", 2701) == -1
    assert bs.transfer(7, "account1", "account2", 200) == 2500

    bs = BankingSystem()
    assert bs.createAccount(1, "acc1") is True
    assert bs.createAccount(2, "acc2") is True
    assert bs.createAccount(3, "acc3") is True
    assert bs.deposit(4, "acc1", 1000) == 1000
    assert bs.deposit(5, "acc2", 500) == 500
    assert bs.deposit(6, "acc3", 0) == 0
    assert bs.transfer(7, "acc1", "acc1", 100) == -1
    assert bs.transfer(8, "acc1", "acc2", 300) == 700
    assert bs.transfer(9, "acc2", "acc3", 800) == 0
    assert bs.transfer(10, "acc1", "acc3", 700) == 0

    bs = BankingSystem()
    assert bs.createAccount(1, "") is True
    assert bs.createAccount(2, "") is False
    assert bs.createAccount(3, "user@123") is True
    assert bs.createAccount(4, "user-account_1") is True
    assert bs.deposit(5, "", 1000000) == 1000000
    assert bs.deposit(6, "user@123", 2500) == 2500
    assert bs.transfer(7, "", "nonexistent", 500) == -1
    assert bs.transfer(8, "nonexistent", "user@123", 100) == -1
    assert bs.transfer(9, "", "user@123", 250000) == 750000
    assert bs.transfer(10, "", "user-account_1", 750000) == 0
    assert bs.deposit(11, "", 50) == 50

    bs = BankingSystem()
    assert bs.createAccount(1, "account3") is True
    assert bs.createAccount(2, "account2") is True
    assert bs.createAccount(3, "account1") is True
    assert bs.deposit(4, "account2", 2000) == 2000
    assert bs.deposit(5, "account3", 3000) == 3000
    assert bs.deposit(6, "account1", 4000) == 4000
    assert bs.topSpenders(7, 3) == ["account1(0)", "account2(0)", "account3(0)"]
    assert bs.transfer(8, "account3", "account2", 500) == 3500
    assert bs.transfer(9, "account3", "account1", 1000) == 2500
    assert bs.transfer(10, "account1", "account2", 2500) == 500
    assert bs.topSpenders(11, 3) == ["account1(2500)", "account3(1500)", "account2(0)"]

    bs = BankingSystem()
    assert bs.createAccount(1, "account1") is True
    assert bs.createAccount(2, "account2") is True
    assert bs.deposit(3, "account1", 2000) == 2000
    assert bs.pay(4, "account1", 1000) == "payment1"
    assert bs.pay(100, "account1", 1000) == "payment2"
    assert bs.getPaymentStatus(101, "non-existing", "payment1") == ""
    assert bs.getPaymentStatus(102, "account2", "payment1") == ""
    assert bs.getPaymentStatus(103, "account1", "payment1") == "IN_PROGRESS"
    assert bs.topSpenders(104, 2) == ["account1(2000)", "account2(0)"]
    assert bs.deposit(86400003, "account1", 100) == 100
    assert bs.getPaymentStatus(86400004, "account1", "payment1") == "CASHBACK_RECEIVED"
    assert bs.deposit(86400005, "account1", 100) == 220
    assert bs.deposit(86400099, "account1", 100) == 320
    assert bs.deposit(86400100, "account1", 100) == 440

    bs = BankingSystem()
    assert bs.createAccount(1, "account1") is True
    assert bs.createAccount(2, "account2") is True
    assert bs.deposit(3, "account1", 2000) == 2000
    assert bs.deposit(4, "account2", 2000) == 2000
    assert bs.pay(5, "account2", 300) == "payment1"
    assert bs.transfer(6, "account1", "account2", 500) == 1500
    assert bs.mergeAccounts(7, "account1", "non-existing") is False
    assert bs.mergeAccounts(8, "account1", "account1") is False
    assert bs.mergeAccounts(9, "account1", "account2") is True
    assert bs.deposit(10, "account1", 100) == 3800
    assert bs.deposit(11, "account2", 100) == -1
    assert bs.getPaymentStatus(12, "account2", "payment1") == ""
    assert bs.getPaymentStatus(13, "account1", "payment1") == "IN_PROGRESS"
    assert bs.getBalance(14, "account2", 1) == -1
    assert bs.getBalance(15, "account2", 9) == -1
    assert bs.getBalance(16, "account1", 11) == 3800
    assert bs.deposit(86400005, "account1", 100) == 3906
