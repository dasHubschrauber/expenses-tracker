package models

type Expense struct {
	Id     int     `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Desc   string  `json:"desc"`
}
