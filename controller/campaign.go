package controller

import (
	"net/http"
	"strconv"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddCampaign godoc
// @Summary AddCampaign api used to create new campaign
// @Description Create a new campaign
// @Tags campaigns
// @MimeType
// @Produce json
// @Param campaign body entity.Campaign true "New Campaign"
// @Success 200 {object} translate.ResultSuccess{data=entity.Campaign} desc
// @ailure 400 {object} translate.ResultErrors
// @Router /campaigns [post]
// @Security ApiKeyAuth
func AddCampaign(c echo.Context) error {
	campaign := new(e.Campaign)
	err := c.Bind(&campaign)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	company, err := r.FindCompany(c)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error)
	}
	campaign.CompanyID = company.ID
	campaign.CreatedBy = utils.GetEmail(c)
	err = r.AddCampaign(campaign)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, campaign)
}

// FindCampaignByID godoc
// @Summary FindcampaignById used to find campaign by it's primary key
// @Description find campaign by id
// @Tags campaigns
// @Accept json
// @Produce json
// @Param id path string true "Campaign ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/{id} [get]
// @Security ApiKeyAuth
func FindCampaignByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := r.FindCampaignByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, campaign)
}

// GetCampaigns godoc
// @Summary GetCampaigns used to find campaign by specific params
// @Description find campaign by date
// @Tags campaigns
// @Accept json
// @Produce json
// @Param filter query repository.CampaignsFilter false "CampaignsFilter"
// @Success 200 {object} translate.ResultSuccess{data=[]entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns [get]
// @Security ApiKeyAuth
func GetCampaigns(c echo.Context) error {
	filter := new(r.CampaignsFilter)
	if err := c.Bind(filter); err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	campaigns := r.GetCampaigns(filter, c)
	return t.Success(c, campaigns)
}

// GetAllSocialMedia godoc
// @Summary GetAllSocialMedia used to find all social media list
// @Description find campaign by date
// @Tags campaigns
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.SocialMedia} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/social-media [get]
// @Security ApiKeyAuth
func GetAllSocialMedia(c echo.Context) error {
	socialMediaList := r.GetAllSocialMedia()
	return t.Success(c, socialMediaList)
}

// GetPaymentTerms godoc
// @Summary GetPaymentTerms used to find all payment terms list
// @Description find all payment terms
// @Tags campaigns
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.PaymentTerms} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/payment-terms [get]
// @Security ApiKeyAuth
func GetPaymentTerms(c echo.Context) error {
	paymentTerms := r.GetPaymentTerms()
	return t.Success(c, paymentTerms)
}

// GetPaymentDays godoc
// @Summary GetPaymentDays used to find all payment days list
// @Description find all payment days
// @Tags campaigns
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.PaymentDays} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/payment-days [get]
// @Security ApiKeyAuth
func GetPaymentDays(c echo.Context) error {
	paymentDays := r.GetPaymentDays()
	return t.Success(c, paymentDays)
}

// GetCampaignInfo godoc
// @Summary GetCampaignInfo used to find campaign information
// @Description used to find campaign information
// @Tags campaigns
// @Accept json
// @Produce json
// @Param id path string true "Campaign ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.CampaignInfo} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/info/{id} [get]
// @Security ApiKeyAuth
func GetCampaignInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	info, err := r.GetCampaignInfo(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, info)
}
