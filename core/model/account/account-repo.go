package account

type AccountRepo interface {
	FindAccount(AccountID) (*Account, error)
	SaveAccount(*Account) error
}
