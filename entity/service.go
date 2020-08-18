package entity

// Service used for talent to write their service
type Service struct {
	ID                 uint         `gorm:"primary_key" json:"id"`
	ServiceDescription string       `json:"description" sql:"index" gorm:"not null" validate:"required"`
	CategoryID         uint         `json:"category_id" gorm:"not null"`
	Category           *Category    `json:"category" swaggerignore:"true"`
	SubCategoryID      uint         `json:"sub_category_id" gorm:"not null" `
	SubCategory        *SubCategory `json:"sub_category" swaggerignore:"true"`
	ImageURL           string       `json:"image_url"`
	TalentID           uint         `json:"talent_id" gorm:"not null"`
	Topic              []Expertise  `json:"topics" gorm:"not null;many2many:service_topics;"`
	Portofilios        []Image      `json:"portfolios" gorm:"many2many:portfolios;"`
	StartPrice         float64      `json:"start_price"`
}
