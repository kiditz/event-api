package repository

import (
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// AddToCart godoc
func AddToCart(cart *e.Cart, c echo.Context) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		existingCart := e.Cart{}
		if !tx.Where("device_id = ?", cart.DeviceID).First(&existingCart).RecordNotFound() {
			if cart.CategoryID != existingCart.CategoryID {
				return fmt.Errorf("cart_category_id_device_id")
			}
		}
		if err := tx.Save(&cart).Error; err != nil {
			return err
		}
		return nil
	})
}

//DeleteCart delete cart by loggedin
func DeleteCart(deviceID string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("device_id = ?", deviceID).Delete(e.Cart{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//GetCarts delete cart by loggedin
func GetCarts(deviceID string) []e.Cart {
	carts := []e.Cart{}
	db.DB.Where("device_id = ?", deviceID).Preload("Service.Category").Preload("Service.SubCategory").Preload("Service.User").Preload("Service.Background").Find(&carts)
	return carts
}

// AddInvitation godoc
func AddInvitation(invitations *[]e.Invitation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		for _, invitation := range *invitations {
			invitation.Status = e.ACTIVE
			if err := tx.Create(&invitation).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

//GetInvitations find invitations by talent logged in
func GetInvitations(email string, limitOffset e.LimitOffset) []e.Invitation {
	if limitOffset.Limit <= 0 {
		limitOffset.Limit = 10
	}
	invitations := []e.Invitation{}
	query := db.DB
	query = query.Joins("JOIN services s ON s.id = invitations.service_id")
	query = query.Joins("JOIN users u ON s.user_id = u.id")
	query = query.Where("u.email = ?", email)
	query = query.Preload("Brief")
	query = query.Preload("Service.Category").Preload("Service.SubCategory").Preload("Service.Topics").Preload("Service.Background").Preload("Service.User")
	query = query.Preload("Brief.Company.Image")
	query = query.Preload("Brief.Company.User")
	query = query.Preload("Brief.Category")
	query = query.Preload("Brief.Location")
	query = query.Preload("Brief.Company")
	query = query.Preload("Brief.PaymentTerms")
	query = query.Order("id desc").Limit(limitOffset.Limit).Offset(limitOffset.Offset).Find(&invitations)
	return invitations
}

// GetCountInvitation godoc
func GetCountInvitation(c echo.Context) e.InvitationCount {
	var invotationCount e.InvitationCount
	query := db.DB
	query = query.Joins("JOIN services s ON s.id = invitations.service_id")
	query = query.Joins("JOIN users u ON u.id = s.user_id AND u.email = ?", utils.GetEmail(c))
	query = query.Where("invitations.status = ?", e.ACTIVE)
	query = query.Find(&e.Invitation{}).Count(&invotationCount.Count)
	return invotationCount
}

// AcceptInvitation godoc
func AcceptInvitation(quote *e.Quotation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var invitation e.Invitation
		if !tx.Set("gorm:query_option", "FOR UPDATE").Where("id = ?", quote.InvitationID).Find(&invitation).RecordNotFound() {
			if invitation.Status == e.ACCEPTED {
				return fmt.Errorf("status_was_accepted")
			}
			// Save Invitation
			invitation.Status = e.ACCEPTED
			if err := tx.Save(&invitation).Error; err != nil {
				return err
			}
		}
		// Save Quotation
		quote.Status = e.ACTIVE
		if err := tx.Save(&quote).Error; err != nil {
			return err
		}
		return nil
	})
}

// RejectInvitation godoc
func RejectInvitation(reject *e.RejectInvitation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var invitation e.Invitation
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("id = ?", reject.InvitationID).Find(&invitation).Error; err != nil {
			return err
		}
		if invitation.Status == e.ACCEPTED {
			return fmt.Errorf("reject_is_not_allowed")
		}
		// Save Invitation
		invitation.Status = e.REJECTED
		if err := tx.Save(&invitation).Error; err != nil {
			return err
		}
		return nil
	})
}

//GetQuotationsList used to find quotations list by campaig id
func GetQuotationsList(filter *e.FilteredQuotations) []e.Quotation {
	quotations := []e.Quotation{}
	query := db.DB
	if filter.Status == e.ACTIVE {
		query = query.Where("brief_id = ? AND status in (?)", filter.BriefID, []string{e.ACTIVE, e.DECLINED, e.APPROVED})
	} else {
		query = query.Where("brief_id = ? AND status = ?", filter.BriefID, filter.Status)
	}
	query = query.Preload("Service.Category")
	query = query.Preload("Service.Background")
	query = query.Preload("Service.SubCategory")
	query = query.Preload("Service.User")
	query = query.Find(&quotations)

	return quotations
}

//GetQuotations used to find quotations list by campaig id
func GetQuotations(filter *e.FilteredQuotations) []e.QuotationList {
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Status == "" {
		filter.Status = e.ACTIVE
	}
	quotations := []e.QuotationList{}
	query := db.DB.Table("quotations q")
	query = query.Select("q.id, q.price, q.message, u.name, u.image_url, q.status, q.created_at, concat(c.name, ' | ', sc.name), '', p.currency")
	query = query.Joins("JOIN services s ON q.service_id = s.id")
	query = query.Joins("JOIN users u ON s.user_id = u.id")
	query = query.Joins("JOIN categories c ON c.id = s.category_id")
	query = query.Joins("JOIN briefs p ON p.id = q.brief_id")
	query = query.Joins("JOIN sub_categories sc ON sc.id = s.sub_category_id")
	if filter.Status == e.ACTIVE {
		query = query.Where("q.brief_id = ? AND q.status in (?)", filter.BriefID, []string{e.ACTIVE, e.DECLINED, e.APPROVED})
	} else {
		query = query.Where("q.brief_id = ? AND q.status = ?", filter.BriefID, filter.Status)
	}
	rows, err := query.Offset(filter.Offset).Limit(filter.Limit).Order("q.id desc").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Printf("Wrong query :%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		quotation := e.QuotationList{}
		err := rows.Scan(&quotation.ID, &quotation.Price, &quotation.Message, &quotation.Name, &quotation.ImageURL, &quotation.Status, &quotation.CreatedAt, &quotation.ServiceCategory, &quotation.ServiceImageURL, &quotation.Currency)
		if err != nil {
			fmt.Println(err)
		}
		quotations = append(quotations, quotation)
	}
	return quotations
}

// ApproveQuotation godoc
func ApproveQuotation(quote *e.QuotationIdentity) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var quotation e.Quotation
		if err := tx.Where("id = ?", quote.QuotationID).Find(&quotation).Error; err != nil {
			return err
		}
		if quotation.Status == e.APPROVED {
			return fmt.Errorf("status_was_approved")
		}
		// Save Quotation
		quotation.Status = e.APPROVED
		if err := tx.Save(&quotation).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeclineQuotation godoc
func DeclineQuotation(quote *e.QuotationIdentity) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var quotation e.Quotation
		if err := tx.Where("id = ?", quote.QuotationID).Find(&quotation).Error; err != nil {
			return err
		}
		if quotation.Status == e.APPROVED {
			return fmt.Errorf("status_was_approved")
		}
		if quotation.Status == e.DECLINED {
			return fmt.Errorf("status_was_declined")
		}
		// Save Quotation
		quotation.Status = e.DECLINED
		if err := tx.Save(&quotation).Error; err != nil {
			return err
		}
		if quotation.InvitationID != 0 {
			var invitation e.Invitation
			if err := tx.Where("id = ?", quotation.InvitationID).Find(&invitation).Error; err != nil {
				return err
			}
			invitation.Status = e.ACTIVE
			if err := tx.Save(&invitation).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

//AddOrder godoc
func AddOrder(c echo.Context, order *e.Order) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.Set("gorm:association_autoupdate", false)
		downPayment, _ := strconv.ParseFloat(order.CustomField2, 64)
		billing, _ := strconv.ParseFloat(order.CustomField1, 64)
		order.TransactionDetails.DownPayment = downPayment
		order.TransactionDetails.Billing = billing

		briefID, _ := strconv.Atoi(order.CustomField3)
		order.BriefID = uint(briefID)
		var brief e.Brief
		tx.Where("id = ?", briefID).Preload("Company").Find(&brief)
		order.UserID = brief.Company.UserID
		if err := tx.Save(&order).Error; err != nil {
			fmt.Println(err.Error())
			return err
		}
		for _, s := range order.ItemDetails {
			s.OrderID = order.ID
			if err := tx.Save(&s).Error; err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
		return nil
	})
}

//AddPayementNotification godoc
func AddPayementNotification(c echo.Context, payment *e.PaymentNotification) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.Set("gorm:association_autoupdate", false)
		serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
		// Handle signature
		baseSignature := fmt.Sprintf("%s%s%s%s", payment.OrderID, payment.StatusCode, payment.GrossAmount, serverKey)
		h := sha512.New()
		h.Write([]byte(baseSignature))
		bs := h.Sum(nil)
		fmt.Printf("Input : %s \n", baseSignature)
		signature := fmt.Sprintf("%x", bs)
		if signature != payment.SignatureKey {
			return fmt.Errorf("Invalid Signature")
		}
		if err := tx.Save(&payment).Error; err != nil {
			fmt.Println(err.Error())
			return err
		}
		if payment.TransactionStatus == txCapture {
			if payment.FraudStatus == fraudChallenge {
				fmt.Println("waiting_payment")
			} else if payment.FraudStatus == fraudAccept {
				return onPaidStatus(tx, payment)
			} else {
				return fmt.Errorf("invalid_order")
			}
		} else if payment.TransactionStatus == txCancel {
			return onCancelStatus(tx, payment)
		} else if payment.TransactionStatus == txDeny {
			return onCancelStatus(tx, payment)
		} else if payment.TransactionStatus == txSettlement {
			return onPaidStatus(tx, payment)
		} else if payment.TransactionStatus == txExpire {
			return onCancelStatus(tx, payment)
		} else {
			return fmt.Errorf("invalid_order")
		}
		if payment.FraudStatus == fraudDeny {
			return onCancelStatus(tx, payment)
		}
		return fmt.Errorf("invalid_order")
	})
}

const (
	txAuthorize  = "authorize"
	txCapture    = "capture"
	txSettlement = "settlement"
	txDeny       = "deny"
	txPending    = "pending"
	txCancel     = "cancel"
	txRefund     = "refund"
	txExpire     = "expire"

	fraudAccept    = "accept"
	fraudDeny      = "deny"
	fraudChallenge = "challenge"
)

func onPaidStatus(tx *gorm.DB, payment *e.PaymentNotification) error {
	return editOrder(tx, payment, e.ACTIVE)
}

func onCancelStatus(tx *gorm.DB, payment *e.PaymentNotification) error {
	return editOrder(tx, payment, e.BOOKING)
}

func editOrder(tx *gorm.DB, payment *e.PaymentNotification, status string) error {
	var order e.Order
	var brief e.Brief
	var billing e.Billing
	var incomes []e.Income
	query := tx
	query = query.Joins("JOIN transaction_details t ON t.id = orders.transaction_detail_id AND t.order_id = ?", payment.OrderID)
	query = query.Preload("TransactionDetails")
	if query.First(&order).RecordNotFound() {
		return fmt.Errorf("order_not_found")
	}
	if tx.First(&brief, order.BriefID).RecordNotFound() {
		return fmt.Errorf("brief_not_found")
	}
	startDate := time.Now().UTC()
	endDate := time.Date(2100, 1, 1, 12, 0, 0, 0, time.UTC)
	if brief.StartDate == nil {
		brief.StartDate = &startDate
	}
	if brief.EndDate == nil {
		brief.EndDate = &endDate
	}
	if brief.Status != e.CLOSED {
		brief.Status = status
		if err := tx.Save(&brief).Error; err != nil {
			return err
		}
	}
	if brief.Status == e.CLOSED {
		if !tx.Where("brief_id = ?", order.BriefID).Find(&billing).RecordNotFound() {
			if billing.Amount == order.TransactionDetails.GrossAmount {
				billing.HasPaid = true
				if err := tx.Save(&billing).Error; err != nil {
					return err
				}
			}
		}
		tx.Where("brief_id = ?", order.BriefID).Find(&incomes)
		if len(incomes) > 0 {
			for _, income := range incomes {
				income.CanWithdrawal = true
				fee, _ := strconv.Atoi(os.Getenv("AGENCY_FEE"))
				income.AgencyFee = fee
				if err := tx.Save(&income).Error; err != nil {
					return err
				}
			}
		}
	}

	layout := "2006-01-02 15:04:05"
	trxTime, err := time.Parse(layout, payment.TransactionTime)
	if err != nil {
		return err
	}
	order.TransactionTime = trxTime
	order.TransactionStatus = payment.TransactionStatus
	if err := tx.Save(&order).Error; err != nil {
		return err
	}

	return nil
}
