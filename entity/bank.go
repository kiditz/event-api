package entity

// Bank godoc
type Bank struct {
	Code string `json:"code" gorm:"primary_key"`
	Name string `json:"name"`
}
