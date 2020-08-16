package entity

import "time"

// Campaign is model for database campaigns
type Campaign struct {
	ID             uint          `gorm:"primary_key" json:"id"`
	Title          string        `json:"title" gorm:"type:varchar(60);index;not null" validate:"required"`
	Detail         string        `json:"detail" validate:"required" gorm:"not null"`
	Currency       string        `json:"currency" validate:"required" gorm:"not null"`
	Location       *Location     `json:"location,omitempty" gorm:"foreignkey:LocationID"`
	LocationID     uint          `json:"location_id" gorm:"index;"`
	PaymentTerms   *PaymentTerms `json:"payment_terms,omitempty" gorm:"foreignkey:PaymentTermsID"`
	PaymentTermsID uint          `json:"payment_terms_id"`
	PaymentDaysID  uint          `json:"payment_days_id"`
	PaymentDays    *PaymentDays  `json:"payment_days,omitempty" gorm:"foreignkey:PaymentDaysID"`
	StartDate      *time.Time    `json:"start_date"`
	EndDate        *time.Time    `json:"end_date" validate:"gtefield=StartDate"`
	StartPrice     float64       `json:"start_price" gorm:"not null" validate:"gte=50000.0,required"`
	EndPrice       float64       `json:"end_price" gorm:"not null" validate:"gtefield=StartPrice,required"`
	StaffAmount    uint          `json:"staff_amount" gorm:"not null" validate:"gte=1,required"`
	Status         string        `json:"status" gorm:"type:varchar(60);not null;default:'booking'"`
	Model
}
