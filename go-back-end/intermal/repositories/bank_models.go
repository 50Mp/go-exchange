package repositories

type BankRepositoryModel struct {
	id       int    `gorm:"primaryKey"`
	index    int    `gorm:"not null"`
	DateTime string `gorm:"not null"`
	Count    int    `gorm:"not null"`
	BankName string `gorm:"not null"`
	Icon     string `gorm:"not null"`
	Currency string `gorm:"not null"`
	Sell     string `gorm:"not null"`
	Buy      string `gorm:"not null"`
}
