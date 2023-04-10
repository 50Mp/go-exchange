package repositories

type BankRepositoryModel struct {
	Id       int `gorm:"primaryKey"`
	DateTime string
	TimeDay  int
	BankName string
	Icon     string
	Currency string
	Sell     uint32
	Buy      uint32
}

type BankRepository interface {
	CreateBank() error
	GetBankCurrency() ([]BankRepositoryModel, error)
}
