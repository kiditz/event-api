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

// ACTIVE activate
const ACTIVE = "active"

// ACCEPTED accepted
const ACCEPTED = "accepted"

// REJECTED reject
const REJECTED = "rejected"

// RUNNING started
const RUNNING = "running"

//ONREVIEW godoc
const ONREVIEW = "on_review"

// CLOSED godoc
const CLOSED = "closed"

// S50_50 godoc
const S50_50 = "closed"
