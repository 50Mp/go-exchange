package repositories

import (
	"gorm.io/gorm"
)

type BankRepositoryDb struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB, baseUrlBCEL, bcelA, bcelB string) BankRepository {
	return BankRepositoryDb{db: db}
}

func (db BankRepositoryDb) CreateBank() error {
	// currency, err := exchangebank.BcelExchange()

	// if err != nil {
	// 	return err
	// }

	// var currencys []BankRepositoryModel

	// for _, c := range currency {

	// 	sell, _ := strconv.ParseUint(c.InSideSell, 10, 32)
	// 	buy, _ := strconv.ParseUint(c.Buy, 10, 32)
	// 	curren := BankRepositoryModel{
	// 		Id:       c.Id,
	// 		DateTime: c.Dateofdate.Date,
	// 		TimeDay:  c.NumberOftime.Id,
	// 		BankName: "",
	// 		Icon:     c.Icon,
	// 		Currency: c.Currency,
	// 		Sell:     uint32(sell),
	// 		Buy:      uint32(buy),
	// 	}
	// 	currencys = append(currencys, curren)
	// }

	// db.db.Create(currencys)
	return nil
}
