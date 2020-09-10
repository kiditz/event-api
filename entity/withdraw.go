package entity

import "time"

// Withdraw godoc
type Withdraw struct {
	ID             uint       `json:"id" gorm:"primary_key"`
	UserBankID     uint       `json:"user_bank_id"`
	Amount         float64    `json:"amount"`
	WithdrawDate   *time.Time `json:"withdraw_date"`
	TransferStatus string     `json:"transfer_status"`
	IncomeID       uint       `json:"income_id"`
}
