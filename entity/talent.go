package entity

import "time"

// Talent account
type Talent struct {
	ID                      uint          `json:"id" gorm:"primary_key"`
	UserID                  uint          `json:"user_id" swaggerignore:"true"`
	User                    *User         `json:"account" swaggerignore:"true"`
	PhoneNumber             string        `json:"phone"`
	Description             string        `json:"description" gorm:"not null;default:''"`
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
	Occupations             []Occupation  `json:"occupations" gorm:"not null;many2many:talent_occupations;"`
	Model
}

// FilteredTalent use to filter talent
type FilteredTalent struct {
	CategoryID    int64  `query:"category_id" json:"category_id"`
	CampaignID    int64  `query:"campaign_id" json:"campaign_id"`
	ExpertiseName string `query:"expertise_name" json:"expertise_name"`
	Limit         int64  `query:"limit" json:"limit"`
	Offset        int64  `query:"offset" json:"offset"`
	Q             string `query:"q" json:"q"`
	SubCategoryID int64  `query:"sub_category_id" json:"sub_category_id"`
}

// TalentResults go doc
type TalentResults struct {
	StartPrice      float64 `json:"start_price" `
	ServiceID       uint    `json:"service_id"`
	TalentID        uint    `json:"talent_id"`
	Name            string  `json:"name"`
	CategoryName    string  `json:"category_name"`
	SubCategoryName string  `json:"sub_category_name"`
	ImageURL        string  `json:"image_url"`
}
