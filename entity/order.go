package entity

import "time"

// Order godoc
type Order struct {
	ID                  uint                `json:"id" gorm:"primary_key"`
	TransactionDetailID uint                `json:"transaction_detail_id" gorm:"not null;index;"`
	TransactionDetails  *TransactionDetails `json:"transaction_details"`
	TransactionTime     time.Time           `json:"transaction_time"`
	ItemDetails         []ItemDetails       `json:"item_details"`
	TransactionStatus   string              `json:"transaction_status" gorm:"not null;varchar(10);" validate:"required"`
	PaymentStatus       string              `json:"payment_status" gorm:"not null;varchar(10);" validate:"required"`
	UserID              uint                `json:"user_id" gorm:"not null;index;" validate:"required"`
	CampaignID          uint                `json:"campaign_id" gorm:"not null;index;" validate:"required"`
}

// TransactionDetails godoc
type TransactionDetails struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	GrossAmount float64 `json:"gross_amount"`
	OrderID     string  `json:"order_id" gorm:"not null;index;varchar(50);"`
}

// ItemDetails godoc
type ItemDetails struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	Price        float64 `json:"price" validate:"required"`
	Quantity     uint    `json:"quantity" validate:"required"`
	Brand        string  `json:"brand"`
	MerchantName string  `json:"merchant_name"`
	Name         string  `json:"name" gorm:"not null;index;varchar(50);" validate:"required"`
	OrderID      uint    `json:"order_id" gorm:"not null;index;"`
}
