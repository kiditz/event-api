package entity

//UserBank godoc
type UserBank struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID      uint   `json:"user_id" `
	User        *User  `json:"user"`
	AccountNo   string `json:"account_no"`
	AccountName string `json:"account_name"`
	BankID      uint   `json:"bank_id"`
	Model
}
