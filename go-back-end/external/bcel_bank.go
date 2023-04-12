package external

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

func BcelExchange() (*[]Item, error) {

	c := colly.NewCollector(
		colly.AllowedDomains(bcelA, bcelB),
	)

	items := []Item{}

	c.OnHTML(".col-sm-9", func(e *colly.HTMLElement) {

		e.ForEach("tbody", func(_ int, b *colly.HTMLElement) {

			b.ForEach("tr", func(_ int, h *colly.HTMLElement) {

				co, _ := strconv.Atoi(e.ChildAttr("option", "value"))
				item := Item{
					Id: h.Index,
					Dateofdate: Dateofdate{
						Id:    e.Index,
						Title: "ວັນທີ",
						Date:  e.ChildAttr("input", "value"),
					},
					Count: Count{
						Id:    e.Index,
						Title: "ຄັ້ງທີ",
						Count: co,
					},
					BankName: BankName{
						Id:     e.Index,
						BkName: "BCEL",
					},
					Icon:       fmt.Sprintf("%v%v", bcelA, h.ChildAttr("img", "src")),
					Currency:   h.ChildText("td:nth-child(3)"),
					InSideSell: h.ChildText("td:nth-child(4)"),
					OutSell:    h.ChildText("td:nth-child(5)"),
					Buy:        h.ChildText("td:nth-child(6)"),
				}

				items = append(items, item)
			})

		})
	})
	// c.OnRequest(func(request *colly.Request) {
	// 	log.Println("Visiting", request.URL.String())
	// })

	c.OnError(func(r *colly.Response, err error) {
		log.Panic("request data in the URL :", r.Request.URL, err.Error())
		return
	})

	c.Visit(baseUrlBCEL)

	if len(items) > 0 {
		err := CreaeJsonFile(items)
		return &items, err
	} else {
		// log.Println("Scrapping curren failed!", items)
		return nil, errors.New("Scrapping curren failed!")
	}

}
