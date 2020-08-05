package entity

//Image is foreign key of campaign
type Image struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	ImageURL string `json:"image_url" validate:"required" gorm:"not null"`
}
