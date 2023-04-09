package main

import (
	exchangebank "exchange/exchange_bank"
	"fmt"
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

	apbEx, _ := exchangebank.ApbExchange(baseUrlApb, apbA, apbB)
	fmt.Println("apb exchange", apbEx)
	fmt.Println("=============================================================")

	bcelex, _ := exchangebank.BcelExchange(baseUrlBCEL, bcelA, bcelB)

	fmt.Println("BCEL exchange", bcelex)

	fmt.Println("=============================================================")

}
