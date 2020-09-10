package entity

//UserBank godoc
type UserBank struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID      uint   `json:"user_id" `
	User        *User  `json:"user"`
	AccountNo   string `json:"account_no" validate:"required"`
	AccountName string `json:"account_name" validate:"required"`
	BankID      uint   `json:"bank_id"`
	Bank        *Bank  `json:"bank"`
	Model
}
