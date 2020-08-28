package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddCart godoc
// @Summary AddCart api used to create new cart for specific email address
// @Description add to cart
// @Tags orders
// @MimeType
// @Produce json
// @Param talent body entity.Cart true "Add To Cart"
// @Success 200 {object} translate.ResultSuccess{data=entity.Cart} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /cart [post]
func AddCart(c echo.Context) error {
	var cart e.Cart
	err := c.Bind(&cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AddToCart(&cart, c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, cart)
}

// DeleteCart godoc
// @Summary DeleteCart api used to delete cart for specific device
// @Description add to cart
// @Tags orders
// @Param device_id query string true "Device ID"
// @MimeType
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=entity.Cart} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /cart [delete]
func DeleteCart(c echo.Context) error {
	deviceID := c.QueryParam("device_id")
	err := r.DeleteCart(deviceID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, map[string]string{"device_id": deviceID})
}

// GetCarts godoc
// @Summary GetCarts api used to find cart for specific device
// @Description find carts
// @Tags orders
// @Param device_id query string true "Device ID"
// @MimeType
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.Cart} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /carts [get]
func GetCarts(c echo.Context) error {
	deviceID := c.QueryParam("device_id")
	carts := r.GetCarts(deviceID)

	return t.Success(c, carts)
}

// AddInvitation godoc
// @Summary AddInvitation api used to create new invitation for talent service
// @Description create new invitation
// @Tags orders
// @MimeType
// @Produce json
// @Param talent body entity.Invitation true "Invitation"
// @Success 200 {object} translate.ResultSuccess{data=entity.Invitation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /invitation [post]
func AddInvitation(c echo.Context) error {
	var invitations []e.Invitation
	err := c.Bind(&invitations)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AddInvitation(&invitations)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, invitations)
}

// GetInvitations godoc
// @Summary GetInvitations api used to invitations by user logged in
// @Description find invitations
// @Tags orders
// @MimeType
// @Produce json
// @Param invitation query entity.LimitOffset false "LimitOffset"
// @Success 200 {object} translate.ResultSuccess{data=[]entity.Invitation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /invitations [get]
// @Security ApiKeyAuth
func GetInvitations(c echo.Context) error {
	var limitOffset e.LimitOffset
	err := c.Bind(&limitOffset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	email := utils.GetEmail(c)
	invitations := r.GetInvitations(email, limitOffset)
	return t.Success(c, invitations)
}

// AcceptInvitation godoc
// @Summary AcceptInvitation api used to accept invitation and generate quote
// @Description accept invitation and generate quote
// @Tags orders
// @MimeType
// @Produce json
// @Param quotation body entity.Quotation true "Quotation"
// @Success 200 {object} translate.ResultSuccess{data=entity.Quotation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /invitation/accept [post]
func AcceptInvitation(c echo.Context) error {
	var quotation e.Quotation
	err := c.Bind(&quotation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AcceptInvitation(&quotation)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, quotation)
}

// RejectInvitation godoc
// @Summary RejectInvitation api used to reject invitation
// @Description reject invitation
// @Tags orders
// @MimeType
// @Produce json
// @Param quotation body entity.RejectInvitation true "RejectInvitation"
// @Success 200 {object} translate.ResultSuccess{data=entity.Invitation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /invitation/reject [post]
func RejectInvitation(c echo.Context) error {
	var reject e.RejectInvitation
	err := c.Bind(&reject)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.RejectInvitation(&reject)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, reject)
}

// GetQuotations godoc
// @Summary GetQuotations get quotations by campaign id
// @Description get quotations by campaign id
// @Tags orders
// @Param filtered quotations query entity.FilteredQuotations true "FilteredQuotations"
// @MimeType
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.QuotationList} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /quotations [get]
// @Security ApiKeyAuth
func GetQuotations(c echo.Context) error {
	var filter e.FilteredQuotations
	err := c.Bind(&filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	quotations := r.GetQuotations(&filter)
	return t.Success(c, quotations)
}

// ApproveQuotation godoc
// @Summary ApproveQuotation api used to approve quotation
// @Description approve quotation
// @Tags orders
// @MimeType
// @Produce json
// @Param quotation body entity.QuotationIdentity true "QuotationIdentity"
// @Success 200 {object} translate.ResultSuccess{data=entity.QuotationIdentity} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /quotation/approved [post]
func ApproveQuotation(c echo.Context) error {
	var quoteID e.QuotationIdentity
	err := c.Bind(&quoteID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.ApproveQuotation(&quoteID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, quoteID)
}

// DeclineQuotation godoc
// @Summary DeclineQuotation api used to decline quotation
// @Description decline quotation
// @Tags orders
// @MimeType
// @Produce json
// @Param quotation body entity.QuotationIdentity true "QuotationIdentity"
// @Success 200 {object} translate.ResultSuccess{data=entity.QuotationIdentity} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /quotation/declined [post]
func DeclineQuotation(c echo.Context) error {
	var quoteID e.QuotationIdentity
	err := c.Bind(&quoteID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.DeclineQuotation(&quoteID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, quoteID)
}

// AddOrder godoc
// @Summary AddOrder api used to add order
// @Description add new order
// @Tags orders
// @MimeType
// @Produce json
// @Param order body entity.Order true "Order"
// @Success 200 {object} translate.ResultSuccess{data=entity.Order} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /order/charge [post]
func AddOrder(c echo.Context) error {
	var order e.Order
	err := c.Bind(&order)
	if err != nil {
		print(err)
		return t.Errors(c, http.StatusBadRequest, err)
	}
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err)
	}
	// campaignID, _ := strconv.Atoi(order.CustomField3)
	now := time.Now().UTC()
	sec := now.Unix()
	order.TransactionDetails.OrderID = fmt.Sprintf("INV/%d/%d", sec, now.Year())

	order.TransactionTime = now
	if order.TransactionDetails.GrossAmount > 0 {
		newOrder, err := getSnapToken(&order)
		print(newOrder.CampaignID)
		if err != nil {
			return t.Errors(c, http.StatusBadRequest, err.Error())
		}
		err = r.AddOrder(c, newOrder)
		if err != nil {
			return t.Errors(c, http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":        order.Token,
			"redirect_url": order.RedirectURL,
		})
	}
	err = r.AddOrder(c, &order)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}

	if err != nil {
		return t.Errors(c, http.StatusInternalServerError, err)
	}
	return t.Success(c, order)
}

func getSnapToken(order *e.Order) (*e.Order, error) {
	serverURL := os.Getenv("MIDTRANS_SERVER_URL")
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	requestBody, _ := json.Marshal(order)
	req, err := http.NewRequest("POST", serverURL+"/snap/v1/transactions", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	fmt.Printf("BODY :%s\n", requestBody)
	req.SetBasicAuth(serverKey, "")
	client := &http.Client{}
	resp, err := client.Do(req)
	bodyText, _ := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println(s)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed_to_call_midtrans")
	}
	if err != nil {
		return nil, err
	}

	fmt.Printf("RES :%s\n", s)
	json.Unmarshal(bodyText, &order)
	return order, nil
}
