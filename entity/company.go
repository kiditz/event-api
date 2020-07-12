package entity

// Company is the employer company for create and sharing job
type Company struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"min=30"`
	Location    string `json:"location" validate:"required"`
	Industry    string `json:"industry" validate:"required"`
	Model
}
