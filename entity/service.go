package entity

// Service used for talent to write their service
type Service struct {
	ID                 uint        `gorm:"primary_key" json:"id"`
	ServiceDescription string      `json:"description" sql:"index" gorm:"not null" validate:"required"`
	CategoryID         uint        `json:"category_id" gorm:"not null"`
	Category           Category    `json:"category"`
	SubCategoryID      uint        `json:"sub_category_id" gorm:"not null"`
	SubCategory        SubCategory `json:"sub_category"`
	Images             []Image     `json:"portfolios" gorm:"many2many:portfolios;"`
}
