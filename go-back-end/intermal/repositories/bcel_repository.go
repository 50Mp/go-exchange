package repositories

type BankRepository interface {
	CreateBankBCEL() error
	GetBankCurrencyDateBCEL(datetime string, name string, count int) ([]BankRepositoryModel, error)
	GetBankCurrencyDatetoDateBCEL(name string, count int, start string, end string) ([]BankRepositoryModel, error)
	GetBankCurrencyBCEL(name string, count int, currency string) ([]BankRepositoryModel, error)
}
