package external

type Item struct {
	Id           int          `json:"id"`
	Dateofdate   Dateofdate   `json:"dateofday"`
	NumberOftime NumberOftime `json:"numberoftime"`
	BankName     BankName     `json:"bankname"`
	Icon         string       `json:"icon"`
	Currency     string       `json:"currency"`
	InSideSell   string       `json:"insidesell"`
	OutSell      string       `json:"outselle"`
	Buy          string       `json:"buy"`
}

type BankName struct {
	Id     int    `json:"id"`
	BkName string `json:"name"`
}

type Dateofdate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type NumberOftime struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Numberoftime string `json:"number"`
}
