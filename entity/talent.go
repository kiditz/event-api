package entity

import "time"

// Talent account
type Talent struct {
	ID                      uint          `json:"id" gorm:"primary_key"`
	UserID                  uint          `json:"user_id" swaggerignore:"true"`
	User                    *User         `json:"account" swaggerignore:"true"`
	PhoneNumber             string        `json:"phone"`
	BusinessTypeID          uint          `json:"business_type_id"`
	BusinessType            *BusinessType `json:"business_type" swaggerignore:"true"`
	ImageID                 uint          `json:"image_id" swaggerignore:"true"`
	Image                   *Image        `json:"image"`
	BackgroundImageID       uint          `json:"background_image_id"`
	BackgroundImage         *Image        `json:"background"`
	Height                  uint          `json:"height" validate:"required"`
	BirthDate               *time.Time    `json:"birth_date" validate:"required"`
	Gender                  string        `json:"gender" validate:"required"`
	Location                *Location     `json:"location"`
	LocationID              uint          `json:"location_id"`
	Services                []Service     `json:"services" swaggerignore:"true"`
	InstagramLink           string        `json:"instagram_link" validate:"required"`
	InstagramFollowersCount uint          `json:"instagram_followers_count"`
	FacebookLink            string        `json:"facebook_link"`
	FacebookFollowersCount  uint          `json:"facebook_followers_count"`
	TwitterLink             string        `json:"twitter_link"`
	TwitterFollowersCount   uint          `json:"twitter_followers_count"`
	YoutubeLink             string        `json:"youtube_link"`
	YoutubeFollowersCount   uint          `json:"youtube_followers_count"`
	Engagement              float64       `json:"engagement"`
	IsVerified              bool          `json:"is_verified"`
	Expertises              []Expertise   `json:"expertises" gorm:"not null;many2many:talent_expertises;"`
	Model
}
