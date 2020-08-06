package controller

import (
	"net/http"
	"strconv"
	"time"

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
	var hasErr, tx = t.ValidateTranslator(c, campaign)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}

	campaign.CreatedBy = utils.GetUser(c)["email"].(string)
	err = r.AddCampaign(campaign)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
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

// GetCampaignByDate godoc
// @Summary GetCampaignByDate used to find campaign by it's start date and end date
// @Description find campaign by date
// @Tags campaigns
// @Accept json
// @Produce json
// @Param date query string true "Date to search"
// @Success 200 {array} translate.ResultSuccess{data=entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/date [get]
// @Security ApiKeyAuth
func GetCampaignByDate(c echo.Context) error {
	dateStr := c.QueryParam("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	campaign, err := r.GetCampaignByDate(date.Format("2006-01-02"))
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, campaign)
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
