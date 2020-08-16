package entity

//Invitation go doc
type Invitation struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	ServiceID  uint   `json:"service_id" gorm:"not null;index;" validate:"required"`
	CampaignID uint   `json:"campaign_id" gorm:"not null;index;" validate:"required"`
	Status     string `json:"status"`
	Model
}
