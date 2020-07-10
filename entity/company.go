package entity

// Company is the employer company for create and sharing job
type Company struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Industry    string `json:"industry"`
	Model
}
