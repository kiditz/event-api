package entity

// Bank godoc
type Bank struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Code string `json:"code"`
	Name string `json:"name"`
}
