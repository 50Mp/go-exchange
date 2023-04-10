package main

import (
	"exchange/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	baseUrlApb = "https://www.apb.com.la/exchange-rate-detail.php"
	apbA       = "www.apb.com.la"
	apbB       = "apb.com.la"

	baseUrlBCEL = "https://www.bcel.com.la/bcel/exchange-rate.html"
	bcelA       = "www.bcel.com.la"
	bcelB       = "bcel.com.la"
)

func main() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=exchange password=123 dbname=exchangedb port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// apbEx, _ := exchangebank.ApbExchange(baseUrlApb, apbA, apbB)
	// fmt.Println("apb exchange", apbEx)
	// fmt.Println("=============================================================")

	// bcelex, _ := exchangebank.BcelExchange(baseUrlBCEL, bcelA, bcelB)

	// fmt.Println("BCEL exchange", bcelex)

	// fmt.Println("=============================================================")

	bcelRepo := repositories.NewBankRepository(db, baseUrlBCEL, bcelA, bcelB)
	err = bcelRepo.CreateBank(baseUrlBCEL, bcelA, bcelB)
	if err != nil {
		panic("failed to create data")
	}

}
