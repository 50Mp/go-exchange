package external

import (
	"encoding/json"
	"log"
	"os"
)

func CreaeJsonFile(items []Item) error {

	jsonlist, err := json.Marshal(items)
	for i:= 0; i <2 ; i++ {
		
	}
	if err != nil {
		log.Panic("Marshal", err)
	}

	err = os.WriteFile("../bcel_exchange.json", jsonlist, 0644)

	if err != nil {
		log.Panic("WriteFile", err)
	}
	return err
}
