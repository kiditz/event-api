package entity

//PaymentDays required for payment tems
type PaymentDays struct {
	ID   uint `gorm:"primary_key" json:"id"`
	Days uint `json:"days" validate:"required"`
}
