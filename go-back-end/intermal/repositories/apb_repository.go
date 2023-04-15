package repositories

type ApbBankRepository interface {
	CreateBankAPB() error
	GetBankCurrencyDateAPB(datetime string, name string, count int) ([]BankRepositoryModel, error)
	GetBankCurrencyDatetoDateAPB(name string, count int, start string, end string) ([]BankRepositoryModel, error)
	GetBankCurrencyAPB(name string, count int, currency string) ([]BankRepositoryModel, error)
}
