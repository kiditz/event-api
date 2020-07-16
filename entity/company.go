package entity

// Company is the employer company for create and sharing job
type Company struct {
	ID          uint     `json:"id" gorm:"primary_key"`
	Name        string   `json:"name" validate:"required"`
	Website     string   `json:"website,omitempty"`
	ImageURL    string   `json:"image_url,omitempty"`
	Description string   `gorm:"type:varchar(300);not null" json:"description" validate:"required,min=30,max=300"`
	Overview    string   `json:"overview"`
	Location    Location `json:"location,omitempty"`
	LocationID  uint     `json:"location_id,omitempty"`
	Industry    string   `json:"industry" validate:"required"`
	Model
}
