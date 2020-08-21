package entity

//Invitation go doc
type Invitation struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	ServiceID  uint      `json:"service_id" gorm:"not null;index;" validate:"required"`
	Service    *Service  `json:"service"`
	CampaignID uint      `json:"campaign_id" gorm:"not null;index;" validate:"required"`
	Campaign   *Campaign `json:"campaign"`
	Status     string    `json:"status"`
	Model
}

// RejectInvitation used to reject invitaion by id
type RejectInvitation struct {
	InvitationID uint `json:"invitation_id" query:"invitation_id"`
}
