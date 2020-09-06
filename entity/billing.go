package entity

import "time"

// Billing godoc
type Billing struct {
	ID      uint       `json:"id" gorm:"primary_key" `
	OrderID string     `json:"order_id"  gorm:"index;" gorm:"not null;"`
	BriefID uint       `json:"brief_id" gorm:"index;" validate:"required"`
	Amount  float64    `json:"amount"`
	UserID  uint       `json:"user_id"`
	User    *User      `json:"user"`
	DueDate *time.Time `json:"due_date"`
}
