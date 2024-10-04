package banksystem

import (
	"bank-system/fixeddeposit"
	"bank-system/useraccount"
	"fmt"
)

type BankSystem struct {
	Users         map[string]*useraccount.UserAccount
	FixedDeposits map[string][]*fixeddeposit.FixedDeposit
}

func NewBankSystem() *BankSystem {
	return &BankSystem{
		Users:         make(map[string]*useraccount.UserAccount),
		FixedDeposits: make(map[string][]*fixeddeposit.FixedDeposit),
	}
}

func (bs *BankSystem) AddUser(user *useraccount.UserAccount) {
	bs.Users[user.AccountNumber] = user
	fmt.Println("New Account Created")
}

func (bs *BankSystem) AddFixedDeposit(fd *fixeddeposit.FixedDeposit) {
	bs.FixedDeposits[fd.AccountNumber] = append(bs.FixedDeposits[fd.AccountNumber], fd)
}

func (bs *BankSystem) DisplayUsers() {
	for _, user := range bs.Users {
		fmt.Printf("Account Number: %s, Holder Name: %s, Age: %d\n", user.AccountNumber, user.Name, user.Age)
	}
}

func (bs *BankSystem) DisplayFixedDeposits() {
	for accNum, fds := range bs.FixedDeposits {
		fmt.Printf("Account Number: %s\n", accNum)
		for _, fd := range fds {
			fmt.Printf("FD Amount: %d, Interest Rate: %.2f%%, Maturity Amount: %.2f\n",
				fd.Amount, fd.InterestRate, fd.MaturityAmount())
		}
		fmt.Println("----------------------------------")
	}
}
