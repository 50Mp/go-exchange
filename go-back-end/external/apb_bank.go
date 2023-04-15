package external

import (
	"errors"
	"fmt"

	"log"

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
				Icon:       fmt.Sprintf("%v/%v", apbA, h.ChildAttr("img", "src")),
				Currency:   h.ChildText("th"),
				InSideSell: h.ChildText("td:nth-child(2)"),
				OutSell:    h.ChildText("td:nth-child(3)"),
				Buy:        h.ChildText("td:nth-child(4)"),
			}

			items = append(items, iten)
		})

	})

	c.OnError(func(r *colly.Response, err error) {
		log.Panic("", r.Request.URL, err.Error())
	})

	c.Visit(baseUrlApb)

	fmt.Println(items)

	if len(items) > 0 {
		err := CreateJsonFile(items, "../apb.json")
		return &items, err
	} else {
		// log.Println("Scrapping curren failed!", items)
		return nil, errors.New("Scrapping curren failed!")
	}
}
