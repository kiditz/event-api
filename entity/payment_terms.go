package entity

// PaymentTerms terms used to campaign payment type
type PaymentTerms struct {
	ID           uint        `gorm:"primary_key" json:"id"`
	Name         string      `json:"name" sql:"index" gorm:"type:varchar(60);not null" validate:"required"`
	Slug         string      `json:"slug" gorm:"type:varchar(20);"`
	PayementDays PaymentDays `json:"payment_days" gorm:"foreignkey:PaymentTermsID"`
}
