package useraccount

type UserAccount struct {
	AccountNumber string
	MobileNumber  string
	Name          string
	Age           int
}

func GenerateAccountNumber(mobileNumber string) string {
	return "BSHDFC" + mobileNumber
}
func NewUserAccount(mobileNumber, name string, age int) *UserAccount {
	return &UserAccount{
		AccountNumber: GenerateAccountNumber(mobileNumber),
		MobileNumber:  mobileNumber,
		Name:          name,
		Age:           age,
	}
}
