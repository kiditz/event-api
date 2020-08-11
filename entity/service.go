package entity

// Service used for talent to write their service
type Service struct {
	ID                 uint        `gorm:"primary_key" json:"id"`
	ServiceDescription string      `json:"description" sql:"index" gorm:"not null" validate:"required"`
	CategoryID         uint        `json:"category_id" gorm:"not null"`
	Category           Category    `json:"category"`
	ImageURL           string      `json:"image_url"`
	SubCategoryID      uint        `json:"sub_category_id" gorm:"not null"`
	TalentID           uint        `json:"talent_id" gorm:"not null"`
	SubCategory        SubCategory `json:"sub_category"`
	ExpertiseID        uint        `json:"expertise_id"`
	Expertise          Expertise   `json:"expertise" gorm:"not null"`
	Portofilios        []Image     `json:"portfolios" gorm:"many2many:portfolios;"`
}
