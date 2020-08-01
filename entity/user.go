package entity

// User for login and logout
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email,omitempty" gorm:"unique;not null"`
	Name     string `json:"name" validate:"required" gorm:"not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Type     string `json:"type,omitempty" gorm:"not null"`
	Model
}
