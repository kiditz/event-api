package entity

//SocialMedia used for social_medias database
type SocialMedia struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" validate:"required" gorm:"type:varchar(60);not null"`
	Icon string `json:"icon" gorm:"type:varchar(20)"`
}
