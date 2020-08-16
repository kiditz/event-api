package entity

//Quotation go doc
type Quotation struct {
	ID         uint    `gorm:"primary_key" json:"id"`
	ServiceID  uint    `json:"service_id" gorm:"not null;index;" validate:"required"`
	CampaignID uint    `json:"campaign_id" gorm:"not null;index" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	Message    string  `json:"message" validate:"required"`
	Status     string  `json:"status"`
	Model
}
