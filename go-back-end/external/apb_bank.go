package external

import (
	"encoding/json"

	"fmt"

	"log"
	"os"

	"github.com/gocolly/colly"
)

func ApbExchange() (*[]Item, error) {

	c := colly.NewCollector(
		colly.AllowedDomains(apbA, apbB),
	)
	var items []Item
	c.OnHTML(".w3_agileits_gain_list", func(

		e *colly.HTMLElement) {

		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
			iten := Item{
				Id: h.Index,
				Dateofdate: Dateofdate{
					Id:   e.Index,
					Date: e.ChildText("h5"),
				},
				Icon:       h.ChildAttr("img", "src"),
				Currency:   h.ChildText("th"),
				InSideSell: h.ChildText("td:nth-child(2)"),
				OutSell:    h.ChildText("td:nth-child(3)"),
				Buy:        h.ChildText("td:nth-child(4)"),
			}

			items = append(items, iten)
		})

	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting", request.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Panic("", r.Request.URL, err.Error())
	})

	c.Visit(baseUrlApb)

	fmt.Println(items)

	jsonlist, err := json.Marshal(items)
	if err != nil {
		log.Panic("Marshal", err)
	}

	err = os.WriteFile("../apb_exchange.json", jsonlist, 0644)

	if err != nil {
		log.Panic("WriteFile", err)
	}

	return &items, nil
}
