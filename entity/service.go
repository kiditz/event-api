package entity

// Service used for talent to write their service
type Service struct {
	ID                 uint         `gorm:"primary_key" json:"id"`
	ServiceTitle       string       `json:"title" sql:"index" gorm:"not null" validate:"required"`
	ServiceDescription string       `json:"description" sql:"index" gorm:"not null" validate:"required"`
	Language           string       `json:"language" sql:"index" gorm:"not null" validate:"required"`
	CategoryID         uint         `json:"category_id" gorm:"not null"`
	Category           *Category    `json:"category"`
	SubCategoryID      uint         `json:"sub_category_id" gorm:"not null" `
	SubCategory        *SubCategory `json:"sub_category"`
	UserID             uint         `json:"user_id" gorm:"not null"`
	User               *User        `json:"user"`
	Topics             []Expertise  `json:"topics" gorm:"not null;many2many:service_topics;"`
	Portfilios         []Portfolio  `json:"portfolios"`
	Background         []Image      `json:"backgrounds" gorm:"many2many:service_backgrounds;"`
	Price              float64      `json:"price"`
	CostPerView        float64      `json:"cpv"`
	Status             string       `json:"status" gorm:"varchar(10)"`
}

// FilteredService use to filter service
type FilteredService struct {
	CategoryID int64 `query:"category_id" json:"category_id"`
}
