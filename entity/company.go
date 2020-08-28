package entity

// Company go doc
type Company struct {
	ID                uint         `gorm:"primary_key" json:"id"`
	UserID            uint         `json:"user_id" swaggerignore:"true"`
	ImageID           uint         `json:"image_id" swaggerignore:"true"`
	Image             *Image       `json:"image"`
	BackgroundImageID uint         `json:"background_image_id" swaggerignore:"true"`
	BackgroundImage   *Image       `json:"background_image"`
	IsUpdated         bool         `json:"is_updated"`
	IsVerified        bool         `json:"is_verified"`
	Name              string       `json:"name" gorm:"type:varchar(60);index;not null" validate:"required"`
	Phone             string       `json:"phone" gorm:"type:varchar(30);index;"`
	Description       string       `json:"description"`
	CategoryID        uint         `json:"category_id"`
	Category          *Category    `json:"category"`
	SubCategoryID     uint         `json:"sub_category_id"`
	SubCategory       *SubCategory `json:"sub_category"`
	Model
}
