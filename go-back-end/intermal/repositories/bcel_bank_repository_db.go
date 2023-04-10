package repositories

import (
	"exchange/external"
	"strconv"

	"gorm.io/gorm"
)

type BankRepositoryDb struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	db.AutoMigrate(&BankRepositoryModel{})

	return BankRepositoryDb{db: db}
}

func (db BankRepositoryDb) CreateBank() error {

	currency, err := external.BcelExchange()

	if err != nil {
		return err
	}

	var count int64

	db.db.Model(&BankRepositoryModel{}).Count(&count)

	if count > 0 {
		return nil
	}

	var currencys []BankRepositoryModel

	for _, c := range currency {

		sell, _ := strconv.ParseUint(c.InSideSell, 10, 32)
		buy, _ := strconv.ParseUint(c.Buy, 10, 32)
		curren := BankRepositoryModel{
			Id:       c.Id,
			DateTime: c.Dateofdate.Date,
			TimeDay:  c.NumberOftime.Id,
			BankName: "",
			Icon:     c.Icon,
			Currency: c.Currency,
			Sell:     uint32(sell),
			Buy:      uint32(buy),
		}
		currencys = append(currencys, curren)
	}

	return db.db.Create(currencys).Error

}

func (db BankRepositoryDb) GetBankCurrency() ([]BankRepositoryModel, error) {

	currency := []BankRepositoryModel{}

	err := db.db.Find(&currency).Error
	return currency, err
}
