package entity

//Quotation godoc
type Quotation struct {
	ID           uint    `gorm:"primary_key" json:"id"`
	ServiceID    uint    `json:"service_id" gorm:"not null;index;" validate:"required"`
	CampaignID   uint    `json:"campaign_id" gorm:"not null;index" validate:"required"`
	InvitationID uint    `json:"invitation_id" gorm:"index"`
	Price        float64 `json:"price" validate:"required"`
	Message      string  `json:"message" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Model
}

//QuotationList godoc
type QuotationList struct {
	ID       uint    `json:"id"`
	Price    float64 `json:"price" `
	Message  string  `json:"message"`
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
	Status   string  `json:"status"`
	Model
}

//FilteredQuotations godoc
type FilteredQuotations struct {
	CampaignID uint   `json:"campaign_id" query:"campaign_id"`
	Status     string `json:"status" query:"status"`
	LimitOffset
}

//QuotationIdentity godoc
type QuotationIdentity struct {
	QuotationID uint `json:"quotation_id" query:"campaign_id"`
}
