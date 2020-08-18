package entity

// Company go doc
type Company struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    uint   `json:"user_id" swaggerignore:"true"`
	ImageID   uint   `json:"image_id" swaggerignore:"true"`
	Image     *Image `json:"image"`
	IsUpdated bool   `json:"is_updated"`
	Name      string `json:"name" gorm:"type:varchar(60);index;not null" validate:"required"`
	Currency  string `json:"currency" validate:"required" gorm:"not null"`
	Country   string `json:"country" validate:"required" gorm:"not null"`
	Model
}
