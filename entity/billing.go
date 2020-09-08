package entity

import "time"

// Billing godoc
type Billing struct {
	ID      uint       `json:"id" gorm:"primary_key" `
	OrderID string     `json:"order_id"  gorm:"index;" gorm:"not null;"`
	BriefID uint       `json:"brief_id" gorm:"index;" validate:"required"`
	Brief   *Brief     `json:"brief,omitempty" swaggerignore:"true"`
	Amount  float64    `json:"amount"`
	UserID  uint       `json:"user_id" gorm:"index;"`
	User    *User      `json:"user"`
	DueDate *time.Time `json:"due_date"`
	HasPaid bool       `json:"has_paid" gorm:"default:'false'"`
}

//BillingFilter godoc
type BillingFilter struct {
	StartDate string `query:"start_date" json:"start_date"`
	EndDate   string `query:"end_date" json:"end_date"`
	Query     string `query:"query" json:"query"`
	Offset    int    `query:"offset" json:"offset"`
	Limit     int    `query:"limit" json:"limit"`
}
