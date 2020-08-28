package entity

// User for login and logout
type User struct {
	ID                 uint   `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Email              string `json:"email,omitempty" gorm:"unique;not null;type:varchar(255);" validate:"required"`
	Name               string `json:"name" gorm:"not null;type:varchar(60);index;"  validate:"required"`
	ImageURL           string `json:"image_url"`
	BackgroundImageURL string `json:"background_image_url"`
	Password           string `json:"-" gorm:"not null;type:varchar(60);" validate:"required"`
	Type               string `json:"type,omitempty" gorm:"not null;type:varchar(20);" validate:"required"`
	Currency           string `json:"currency" validate:"required" gorm:"not null;default:'IDR'"`
	Language           string `json:"language" validate:"required" gorm:"not null;default:'ID'"`
	Model
}

//UserForm godoc
type UserForm struct {
	ID       uint   `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Email    string `json:"email,omitempty" gorm:"unique;not null;type:varchar(255);;" validate:"required"`
	Name     string `json:"name" validate:"required" gorm:"not null;type:varchar(60);" validate:"required"`
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(60);" validate:"required"`
	Type     string `json:"type,omitempty" gorm:"not null;type:varchar(20);" validate:"required"`
	Model
}
