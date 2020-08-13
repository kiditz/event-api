package entity

// Cart used mark service
type Cart struct {
	ID        uint     `json:"id" gorm:"primary_key" `
	IPAddress string   `json:"ip_address" gorm:"ip_address;not null;"`
	TalentID  uint     `json:"talent_id" gorm:"not null"`
	Talent    *Talent  `json:"talent" swaggerignore:"true"`
	ServiceID uint     `json:"service_id" gorm:"not null"`
	Service   *Service `json:"service" swaggerignore:"true"`
}
