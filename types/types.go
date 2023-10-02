package types

type Album struct {
	ID     int64   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
