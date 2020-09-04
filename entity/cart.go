package entity

// Cart used mark service
type Cart struct {
	ID            uint   `json:"id" gorm:"primary_key" `
	DeviceID      string `json:"device_id" gorm:"not null;"`
	DeviceName    string `json:"device_name" gorm:"not null;"`
	DeviceVersion string `json:"device_version" gorm:"not null;"`
	ServiceID     uint   `json:"service_id" gorm:"not null"`
	CategoryID    uint   `json:"category_id" gorm:"not null"`
	// Service       *Service `json:"service" swaggerignore:"true"`
}
