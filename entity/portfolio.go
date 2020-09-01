package entity

// Portfolio godoc
type Portfolio struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	ServiceID uint   `json:"service_id"`
	Model
}
