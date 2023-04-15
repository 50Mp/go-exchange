package external

import (
	"encoding/json"
	"log"
	"os"
)

func CreateJsonFile(items []Item, jsomName string) error {

	jsonlist, err := json.Marshal(items)
	for i := 0; i < 2; i++ {

	}
	if err != nil {
		log.Panic("Marshal", err)
	}

	err = os.WriteFile(jsomName, jsonlist, 0644)

	if err != nil {
		log.Panic("WriteFile", err)
	}
	return err
}
