package entity

import "time"

// Talent account
type Talent struct {
	ID                      uint          `json:"id" gorm:"primary_key" swaggerignore:"true"`
	UserID                  uint          `json:"user_id"`
	ImageID                 uint          `json:"image_id" swaggerignore:"true"`
	BusinessTypeID          uint          `json:"business_type_id" swaggerignore:"true"`
	BusinessType            *BusinessType `json:"business_type"`
	Image                   *Image        `json:"image" validate:"required" gorm:"not null"`
	Height                  uint32        `json:"height" validate:"required"`
	BirthDate               *time.Time    `json:"birth_date" validate:"required"`
	Gender                  string        `json:"gender" validate:"required"`
	Location                *Location     `json:"location"`
	InstagramLink           string        `json:"instagram_link" validate:"required"`
	InstagramFollowersCount int           `json:"instagram_followers_count"`
	FacebookLink            string        `json:"facebook_link"`
	FacebookFollowersCount  int           `json:"facebook_followers_count"`
	TwitterLink             string        `json:"twitter_link"`
	TwitterFollowersCount   int           `json:"twitter_followers_count"`
	YoutubeLink             string        `json:"youtube_link"`
	YoutubeFollowersCount   int           `json:"youtube_followers_count"`
	Engagement              float64       `json:"engagement"`
	IsVerified              bool          `json:"is_verified"`
	Model
}
