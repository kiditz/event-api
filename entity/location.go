package entity

// Location used to save data of user location
type Location struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Country string `json:"country" gorm:"type:varchar(10);not null;index:country" validate:"required"`
	City    string `json:"city" gorm:"type:varchar(10);not null;index:city" validate:"required"`
}
