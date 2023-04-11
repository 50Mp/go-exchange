package main

import (
	"exchange/external"
	"fmt"
)

func main() {

	// db, err := config.Initdb()

	// if err != nil {
	// 	panic("db connect failed")
	// }

	// bcelRepo := repositories.NewBankRepository(db)

	// err = bcelRepo.CreateBank()

	// if err != nil {
	// 	log.Fatalln("Create data from API fialed", err)

	// }

	moneyExchange, err := external.BcelExchange()

	if err != nil {
		panic("Get data failed!")
	}

	fmt.Println(moneyExchange)
}
