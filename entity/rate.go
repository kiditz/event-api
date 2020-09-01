package entity

// Rate used mark service
type Rate struct {
	ID              uint     `json:"id" gorm:"primary_key" `
	Criteria        string   `json:"criteria" gorm:"not null;"`
	Price           float64  `json:"price"`
	Height          uint     `json:"height"`
	TotalFollowers  uint     `json:"total_followers" gorm:"not null;default:'0'"`
	SubCategorySlug string   `json:"sub_category_slug" gorm:"not null" `
	Service         *Service `json:"service" swaggerignore:"true"`
}
