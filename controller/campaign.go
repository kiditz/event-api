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
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns [post]
// @Security ApiKeyAuth
func AddCampaign(c echo.Context) error {
	campaign := new(e.Campaign)
	err := c.Bind(&campaign)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	campaign.CreatedBy = utils.GetUser(c)["email"].(string)
	err = r.AddCampaign(campaign)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, campaign)
}

// FindcampaignByID godoc
// @Summary FindcampaignById used to find campaign by it's primary key
// @Description find campaign by id
// @Tags campaigns
// @Accept json
// @Produce json
// @Param id path string true "Campaign ID"
// @Success 200 {array} translate.ResultSuccess{data=entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/{id} [get]
// @Security ApiKeyAuth
func FindcampaignByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := r.FindCampaignByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, campaign)
}

// GetCampaigns godoc
// @Summary GetCampaigns used to find campaign by it's start date and end date
// @Description find campaign by date
// @Tags campaigns
// @Accept json
// @Produce json
// @Param date query string false "date"
// @Param q query string false "title"
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {array} translate.ResultSuccess{data=entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns [get]
// @Security ApiKeyAuth
func GetCampaigns(c echo.Context) error {
	filter := new(r.CampaignsFilter)
	if err := c.Bind(filter); err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	campaigns := r.GetCampaigns(filter)
	return t.Success(c, campaigns)
}

// GetAllSocialMedia godoc
// @Summary GetAllSocialMedia used to find all social media list
// @Description find campaign by date
// @Tags campaigns
// @Accept json
// @Produce json
// @Success 200 {array} translate.ResultSuccess{data=entity.SocialMedia} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/social-media [get]
// @Security ApiKeyAuth
func GetAllSocialMedia(c echo.Context) error {
	socialMediaList, err := r.GetAllSocialMedia()
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, socialMediaList)
}
