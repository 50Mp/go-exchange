package external

type Name string

const (
	bcel Name = "BCEL"
	ldb  Name = "LDB"
	jdb  Name = "JDB"
	kpv  Name = "KPV"
)

type Item struct {
	Id         int        `json:"id"`
	Dateofdate Dateofdate `json:"dateofday"`
	Count      Count      `json:"count"`
	BankName   BankName   `json:"bankname"`
	Icon       string     `json:"icon"`
	Currency   string     `json:"currency"`
	InSideSell string     `json:"insidesell"`
	OutSell    string     `json:"outselle"`
	Buy        string     `json:"buy"`
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

type Count struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}


