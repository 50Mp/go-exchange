package repositories

import (
	"exchange/external"

	"gorm.io/gorm"
)

type BcelRepositoryDb struct {
	db *gorm.DB
}

func NewBankRepositoryBCEL(db *gorm.DB) BankRepository {

	return BcelRepositoryDb{db: db}
}

func (db BcelRepositoryDb) CreateBankBCEL() error {

	err := db.db.AutoMigrate(&[]BankRepositoryModel{})

	if err != nil {
		return err
	}

	currency, err := external.NewBCEL().BcelExchange()

	if err != nil {
		return err
	}

	currencys := []BankRepositoryModel{}

	for index, c := range *currency {

		curren := BankRepositoryModel{
			id:       index,
			index:    c.Id,
			DateTime: c.Dateofdate.Date,
			Count:    c.Count.Id,
			BankName: c.BankName.BkName,
			Icon:     c.Icon,
			Currency: c.Currency,
			Sell:     c.OutSell,
			Buy:      c.Buy,
		}
		currencys = append(currencys, curren)
	}
	// db.db.Select("count").Find(&currencys)

	// fmt.Println(count)
	db.db.Create(&currencys)

	return nil
}

//Get All By name

// Get Data time and Name
func (db BcelRepositoryDb) GetBankCurrencyDateBCEL(datetime string, name string, count int) ([]BankRepositoryModel, error) {

	currencys := []BankRepositoryModel{}
	date := BankRepositoryModel{
		DateTime: datetime,
		BankName: name,
		Count:    count,
	}

	// currency := db.db.Find(&currencys)
	tx := db.db.Where(date).Find(&currencys)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return currencys, nil
}

//Query scop date start to end

func (db BcelRepositoryDb) GetBankCurrencyDatetoDateBCEL(name string, count int, start string, end string) ([]BankRepositoryModel, error) {

	currencys := []BankRepositoryModel{}
	tx := db.db.Where("date_time <= ? OR date_time >= ? AND bank_name = ? AND count = ? ", start, end, name, count).Find(&currencys)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return currencys, nil

	//	SELECT *
	//
	// FROM bank_repository_models
	// WHERE
	//
	//	date_time <= '2023-04-12'
	//	OR date_time >= '2023-04-7'
	//	AND bank_name = 'BCEL'
	//	AND count = 0;
}

//Get Currency

func (db BcelRepositoryDb) GetBankCurrencyBCEL(name string, count int, currency string) ([]BankRepositoryModel, error) {
	currencys := []BankRepositoryModel{}

	date := BankRepositoryModel{

		BankName: name,
		Count:    count,
		Currency: currency,
	}
	db.db.Where(date).Find(&currencys)

	return currencys, nil
}
