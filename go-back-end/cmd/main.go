package main

import (
	"exchange/intermal/config"
	"exchange/intermal/repositories"
	"fmt"
)

func main() {

	db, err := config.Initdb()

	bcelRepo := repositories.NewBankRepository(db)
	err = bcelRepo.CreateBank()
	if err != nil {
		panic("failed to create data")
	}
	excange, err := bcelRepo.GetBankCurrency()

	if err != nil {
		panic("Get data failed!")
	}

	fmt.Println("", excange)

	// _, err := external.BcelExchange()

	// if err != nil {
	// 	panic(err)
	// }

}
