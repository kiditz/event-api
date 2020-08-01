package entity

// User for login and logout
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email,omitempty" gorm:"unique"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password,omitempty"`
	Model
}
