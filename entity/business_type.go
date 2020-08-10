package entity

// BusinessType is model for database campaigns
type BusinessType struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug" gorm:"not null"`
	Model
}
