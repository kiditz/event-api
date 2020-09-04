package entity

//Quotation godoc
type Quotation struct {
	ID           uint     `gorm:"primary_key" json:"id"`
	ServiceID    uint     `json:"service_id" gorm:"not null;index;" validate:"required"`
	Service      *Service `json:"service" validate:"required"`
	BriefID      uint     `json:"brief_id" gorm:"index" validate:"required"`
	InvitationID uint     `json:"invitation_id" gorm:"index"`
	Price        float64  `json:"price" validate:"required"`
	Message      string   `json:"message" validate:"required"`
	Status       string   `json:"status" validate:"required"`
	Model
}

//QuotationList godoc
type QuotationList struct {
	ID              uint    `json:"id"`
	Price           float64 `json:"price" `
	Currency        string  `json:"currency" `
	Message         string  `json:"message"`
	Name            string  `json:"name"`
	ImageURL        string  `json:"image_url"`
	Status          string  `json:"status"`
	ServiceCategory string  `json:"service_category"`
	ServiceImageURL string  `json:"service_image_url"`
	Model
}

//FilteredQuotations godoc
type FilteredQuotations struct {
	BriefID uint   `json:"brief_id" query:"brief_id"`
	Status  string `json:"status" query:"status"`
	LimitOffset
}

//QuotationIdentity godoc
type QuotationIdentity struct {
	QuotationID uint `json:"quotation_id" query:"quotation_id"`
}
