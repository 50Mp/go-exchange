package modelbank

type Item struct {
	Id         int        `json:"id"`
	Dateofdate Dateofdate `json:"dateofday"`
	Icon       string     `json:"icon"`
	Currency   string     `json:"currency"`
	InSideSell string     `json:"insidesell"`
	OutSell    string     `json:"outselle"`
	Buy        string     `json:"buy"`
}

type Dateofdate struct {
	Id   int    `json:"id"`
	Date string `json:"date"`
}
