package entity

//PaymentNotification godoc
// TODO: Implements this
type PaymentNotification struct {
	ID                uint   `gorm:"primary_key" json:"id"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	SettlementTime    string `json:"settlement_time"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	CustomField3      string `json:"custom_field3"`
	CustomField2      string `json:"custom_field2"`
	CustomField1      string `json:"custom_field1"`
	Currency          string `json:"currency"`
	ApprovalCode      string `json:"approval_code"`
}
