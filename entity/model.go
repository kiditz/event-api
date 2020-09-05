package entity

import "time"

// Model definition
type Model struct {
	CreatedAt time.Time  `json:"created_at" swaggerignore:"true" gorm:"not null;default:now()"`
	CreatedBy string     `json:"created_by,omitempty" swaggerignore:"true"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
	UpdatedBy string     `json:"updated_by,omitempty" swaggerignore:"true"`
}

//BeforeCreate docs
func (u *Model) BeforeCreate() (err error) {
	u.CreatedAt = time.Now().UTC()
	return
}

//BeforeSave docs
func (u *Model) BeforeSave() (err error) {
	date := time.Now().UTC()
	u.UpdatedAt = &date
	return
}

// APPROVED approve
const APPROVED = "approved"

// DECLINED decline
const DECLINED = "declined"

// ACTIVE godocr
const ACTIVE = "active"

// ACCEPTED godoc
const ACCEPTED = "accepted"

// REJECTED godoc
const REJECTED = "rejected"

// RUNNING godoc
const RUNNING = "running"

// BOOKING godoc
const BOOKING = "booking"

//ONREVIEW godoc
const ONREVIEW = "on_review"

// CLOSED godoc
const CLOSED = "closed"

// S50_50 godoc
const S50_50 = "50_50"
