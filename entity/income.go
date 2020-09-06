package entity

import "time"

// Income godoc
type Income struct {
	ID            uint       `json:"id" gorm:"primary_key" `
	OrderID       string     `json:"order_id"  gorm:"index;" gorm:"not null;"`
	BriefID       uint       `json:"brief_id" gorm:"index;" validate:"required"`
	Brief         *Brief     `json:"brief,omitempty" swaggerignore:"true"`
	Amount        float64    `json:"amount"`
	UserID        uint       `json:"user_id"`
	User          *User      `json:"user" swaggerignore:"true"`
	WidrawalDate  *time.Time `json:"withdrawal_time"`
	CanWithdrawal bool       `json:"can_withdrawal" gorm:"not null;default:'false'"`
	HasWithdraw   bool       `json:"has_withdraw" gorm:"not null;default:'false'"`
	Model
}

//IncomeFilter used to filter income
type IncomeFilter struct {
	StartDate     string `query:"start_date" json:"start_date"`
	EndDate       string `query:"end_date" json:"end_date"`
	CanWithdrawal bool   `query:"can_withdrawal" json:"can_withdrawal"`
	HasWithdraw   bool   `query:"has_withdraw" json:"has_withdraw"`
	Offset        int    `query:"offset" json:"offset"`
	Limit         int    `query:"limit" json:"limit"`
}
