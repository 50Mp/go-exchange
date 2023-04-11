package repositories

import (
	"exchange/external"

	"gorm.io/gorm"
)

type BankRepositoryDb struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {

	return BankRepositoryDb{db: db}
}

func (db BankRepositoryDb) CreateBank() error {

	err := db.db.AutoMigrate(&[]BankRepositoryModel{})

	if err != nil {
		return err
	}
	currency, err := external.BcelExchange()

	if err != nil {
		return err
	}
	currencys := []BankRepositoryModel{}

	for index, c := range *currency {

		curren := BankRepositoryModel{
			id:       index,
			index:    c.Id,
			DateTime: c.Dateofdate.Date,
			TimeDay:  c.NumberOftime.Id,
			BankName: c.BankName.BkName,
			Icon:     c.Icon,
			Currency: c.Currency,
			Sell:     c.OutSell,
			Buy:      c.Buy,
		}
		currencys = append(currencys, curren)

	}

	return db.db.Create(&currencys).Error

}

func (db BankRepositoryDb) GetBankCurrency() ([]BankRepositoryModel, error) {

	currency := []BankRepositoryModel{}

	result := db.db.Find(&currency)

	return currency, result.Error
}
