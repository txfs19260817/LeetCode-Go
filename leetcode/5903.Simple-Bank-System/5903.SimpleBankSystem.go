package _903_Simple_Bank_System

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.Withdraw(account1, money) {
		return false
	}
	if !this.Deposit(account2, money) {
		this.balance[account1-1] += money
		return false
	}
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account-1 >= len(this.balance) {
		return false
	}
	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account-1 >= len(this.balance) || this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}
