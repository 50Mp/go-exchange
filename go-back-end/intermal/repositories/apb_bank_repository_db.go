package repositories

import "gorm.io/gorm"

type APBRepository struct {
	db *gorm.DB
}

func NewApbRepositoryAPB(db *gorm.DB) ApbBankRepository {
	return APBRepository{db: db}
}

func (a APBRepository) CreateBankAPB() error {
	return nil
}
func (a APBRepository) GetBankCurrencyDateAPB(datetime string, name string, count int) ([]BankRepositoryModel, error) {
	return nil, nil
}

func (a APBRepository) GetBankCurrencyDatetoDateAPB(name string, count int, start string, end string) ([]BankRepositoryModel, error) {
	return nil, nil
}

func (a APBRepository) GetBankCurrencyAPB(name string, count int, currency string) ([]BankRepositoryModel, error) {
	return nil, nil
}
