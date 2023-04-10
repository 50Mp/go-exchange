package external

import (
	"encoding/json"

	"log"
	"os"

	"github.com/gocolly/colly"
)

func BcelExchange() ([]Item, error) {

	c := colly.NewCollector(
		colly.AllowedDomains(bcelA, bcelB),
	)
	var items []Item

	c.OnHTML(".col-sm-9", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {

			item := Item{
				Id: h.Index,
				Dateofdate: Dateofdate{
					Id:   e.Index,
					Date: e.ChildAttr("input", "value"),
				},
				NumberOftime: NumberOftime{
					Id: e.Index,

					Numberoftime: e.ChildAttr("option", "value"),
				},
				Icon:       h.ChildAttr("img", "src"),
				Currency:   h.ChildText("td:nth-child(3)"),
				InSideSell: h.ChildText("td:nth-child(4)"),
				OutSell:    h.ChildText("td:nth-child(5)"),
				Buy:        h.ChildText("td:nth-child(6)"),
			}

			items = append(items, item)
		})
	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting", request.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Panic("err :", r.Request.URL, err.Error())
		return
	})

	c.Visit(baseUrlBCEL)

	if len(items) > 0 {

		jsonlist, err := json.Marshal(items)
		if err != nil {
			log.Panic("Marshal", err)
		}

		err = os.WriteFile("../bcel_exchange.json", jsonlist, 0644)

		if err != nil {
			log.Panic("WriteFile", err)
		}

		return items, nil
	} else {
		log.Println("Scrapping curren failed!", items)
	}
	return nil, nil
}
