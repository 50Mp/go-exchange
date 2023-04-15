package main

import (
	"exchange/intermal/config"
	"exchange/intermal/repositories"
	"fmt"
	"log"
)

func main() {

	db, err := config.Initdb()

	if err != nil {
		panic("db connect failed")
	}

	bcelRepo := repositories.NewBankRepositoryBCEL(db)

	// err = bcelRepo.CreateBank()

	// if err != nil {
	// 	log.Fatalln("Create data from API fialed", err)

	// }

	fmt.Println("insert sussessfull")

	exchaneg, err := bcelRepo.GetBankCurrencyBCEL("LDB", 1, "USD")
	if err != nil {
		log.Fatalln("Sellect currency fialed", err)

	}

	fmt.Println(exchaneg)

	// moneyExchange, err := external.BcelExchange()

	// if err != nil {
	// 	log.Println("Get data failed!")
	// }

	// for i := 0; i < time.Now().Second(); i++ {

	// 	time.Sleep(time.Second)
	// 	fmt.Println(moneyExchange)
	// 	fmt.Println("=================================================", i)
	// 	fmt.Println(time.Now().Second())

	// }

}

// 2008-01-21
