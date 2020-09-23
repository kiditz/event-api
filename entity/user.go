package entity

// User for login and logout
type User struct {
	ID                 uint      `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Email              string    `json:"email,omitempty" gorm:"unique;not null;type:varchar(255);" validate:"required"`
	Name               string    `json:"name" gorm:"not null;type:varchar(60);index;"  validate:"required"`
	ImageURL           string    `json:"image_url"`
	BackgroundImageURL string    `json:"background_image_url"`
	Services           []Service `json:"services"`
	Active             bool      `json:"active"`
	Password           string    `json:"-" gorm:"not null;type:varchar(60);" validate:"required"`
	Type               string    `json:"type,omitempty" gorm:"not null;type:varchar(20);" validate:"required"`
	Currency           string    `json:"currency" validate:"required" gorm:"not null;default:'IDR'"`
	Language           string    `json:"language" validate:"required" gorm:"not null;default:'ID'"`
	Talent             *Talent   `json:"talent,omitempty"`
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

// FilteredUsers use to filter talent
type FilteredUsers struct {
	CategoryID    int64  `query:"category_id" json:"category_id"`
	BriefID       int64  `query:"brief_id" json:"brief_id"`
	ExpertiseName string `query:"expertise_name" json:"expertise_name"`
	Limit         int64  `query:"limit" json:"limit"`
	Offset        int64  `query:"offset" json:"offset"`
	Query         string `query:"q" json:"q"`
	SubCategoryID int64  `query:"sub_category_id" json:"sub_category_id"`
}
