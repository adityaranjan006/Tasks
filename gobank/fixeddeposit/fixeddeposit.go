package fixeddeposit

import (
	"bank-system/useraccount"
)

type FixedDeposit struct {
	AccountNumber string
	Amount        int64
	Duration      int16
	InterestRate  float64
}

func NewFixedDeposit(user *useraccount.UserAccount, amount int64, duration int16) *FixedDeposit {
	interestRate := calculateInterestRate(duration)
	return &FixedDeposit{
		AccountNumber: user.AccountNumber,
		Amount:        amount,
		Duration:      duration,
		InterestRate:  interestRate,
	}
}

func calculateInterestRate(duration int16) float64 {
	if duration <= 3 {
		return 7.0
	} else if duration <= 5 {
		return 7.2
	} else if duration <= 10 {
		return 7.5
	}
	return 7.9
}

func (fd *FixedDeposit) MaturityAmount() float64 {
	return float64(fd.Amount) + (float64(fd.Amount) * fd.InterestRate * float64(fd.Duration) / 100)
}
