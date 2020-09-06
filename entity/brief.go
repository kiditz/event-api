package entity

import "time"

// Brief is model for database briefs
type Brief struct {
	ID             uint          `gorm:"primary_key" json:"id"`
	CompanyID      uint          `json:"company_id"`
	Company        *Company      `json:"company"`
	Title          string        `json:"title" gorm:"type:varchar(60);index;not null" validate:"required"`
	CategoryID     uint          `json:"category_id" gorm:"not null"`
	Category       *Category     `json:"category"`
	SubCategories  []SubCategory `json:"sub_categories" gorm:"not null;many2many:brief_sub_categories;"`
	Height         uint          `json:"height" gorm:"not null"`
	Gender         string        `json:"gender" gorm:"not null"`
	Detail         string        `json:"detail" gorm:"not null" validate:"required"`
	Currency       string        `json:"currency" validate:"required" gorm:"not null"`
	Location       *Location     `json:"location,omitempty" gorm:"foreignkey:LocationID"`
	LocationID     uint          `json:"location_id" gorm:"index;"`
	PaymentTerms   *PaymentTerms `json:"payment_terms,omitempty" gorm:"foreignkey:PaymentTermsID"`
	PaymentTermsID uint          `json:"payment_terms_id"`
	PaymentDaysID  uint          `json:"payment_days_id"`
	PaymentDays    *PaymentDays  `json:"payment_days,omitempty" gorm:"foreignkey:PaymentDaysID"`
	StartDate      *time.Time    `json:"start_date"`
	EndDate        *time.Time    `json:"end_date" validate:"gtefield=StartDate"`
	StartTime      *time.Time    `json:"start_time"`
	EndTime        *time.Time    `json:"end_time" validate:"gtefield=StartTime"`
	Price          float64       `json:"price" gorm:"not null;default:'100000'" validate:"gte=100000.0,required"`
	StaffAmount    uint          `json:"staff_amount" gorm:"not null" validate:"gte=1,required"`
	Status         string        `json:"status" gorm:"type:varchar(60);not null;default:'booking'"`
	Model
}

//BriefInfo show campaign info
type BriefInfo struct {
	StaffAmount    uint     `json:"staff_amount"`
	ApprovedCount  uint     `json:"approved_count"`
	QuotationCount uint     `json:"quotation_count"`
	Images         []string `json:"images"`
}

//StopBrief show campaign info
type StopBrief struct {
	ID uint `json:"id" query:"id"`
}

//BriefsFilter godoc
type BriefsFilter struct {
	Date   string `query:"date" json:"date"`
	Query  string `query:"q" json:"q"`
	Offset int    `query:"offset" json:"offset"`
	Limit  int    `query:"limit" json:"limit"`
	OnlyMe bool   `query:"onlyme" json:"onlyme"`
}
