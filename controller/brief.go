package controller

import (
	"net/http"
	"strconv"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddBrief godoc
// @Summary AddBrief api used to create new campaign
// @Description Create a new campaign
// @Tags briefs
// @MimeType
// @Produce json
// @Param campaign body entity.Brief true "New Brief"
// @Success 200 {object} translate.ResultSuccess{data=entity.Brief} desc
// @ailure 400 {object} translate.ResultErrors
// @Router /briefs [post]
// @Security ApiKeyAuth
func AddBrief(c echo.Context) error {
	brief := new(e.Brief)
	err := c.Bind(&brief)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = r.AddBrief(c, brief)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, brief)
}

// StopBrief godoc
// @Summary AddBrief api used to create new campaign
// @Description Create a new campaign
// @Tags briefs
// @MimeType
// @Produce json
// @Param campaign body entity.StopBrief true "New Brief"
// @Success 200 {object} translate.ResultSuccess{data=entity.Brief} desc
// @ailure 400 {object} translate.ResultErrors
// @Router /briefs/stop [post]
// @Security ApiKeyAuth
func StopBrief(c echo.Context) error {
	stop := new(e.StopBrief)
	err := c.Bind(&stop)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = r.StopBrief(c, stop.ID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, stop)
}

// FindBriefByID godoc
// @Summary FindcampaignById used to find campaign by it's primary key
// @Description find campaign by id
// @Tags briefs
// @Accept json
// @Produce json
// @Param id path string true "Brief ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.Brief} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /briefs/{id} [get]
// @Security ApiKeyAuth
func FindBriefByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := r.FindBriefByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, campaign)
}

// GetBriefs godoc
// @Summary GetBriefs used to find campaign by specific params
// @Description find campaign by date
// @Tags briefs
// @Accept json
// @Produce json
// @Param filter query entity.BriefsFilter false "BriefsFilter"
// @Success 200 {object} translate.ResultSuccess{data=[]entity.Brief} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /briefs [get]
// @Security ApiKeyAuth
func GetBriefs(c echo.Context) error {
	filter := new(e.BriefsFilter)
	if err := c.Bind(filter); err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	briefs := r.GetBriefs(filter, c)
	return t.Success(c, briefs)
}

// GetAllSocialMedia godoc
// @Summary GetAllSocialMedia used to find all social media list
// @Description find campaign by date
// @Tags briefs
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.SocialMedia} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /briefs/social-media [get]
// @Security ApiKeyAuth
func GetAllSocialMedia(c echo.Context) error {
	socialMediaList := r.GetAllSocialMedia()
	return t.Success(c, socialMediaList)
}

// GetPaymentTerms godoc
// @Summary GetPaymentTerms used to find all payment terms list
// @Description find all payment terms
// @Tags briefs
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.PaymentTerms} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /briefs/payment-terms [get]
// @Security ApiKeyAuth
func GetPaymentTerms(c echo.Context) error {
	paymentTerms := r.GetPaymentTerms()
	return t.Success(c, paymentTerms)
}

// GetPaymentDays godoc
// @Summary GetPaymentDays used to find all payment days list
// @Description find all payment days
// @Tags briefs
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.PaymentDays} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /briefs/payment-days [get]
// @Security ApiKeyAuth
func GetPaymentDays(c echo.Context) error {
	paymentDays := r.GetPaymentDays()
	return t.Success(c, paymentDays)
}

// GetBriefInfo godoc
// @Summary GetBriefInfo used to find campaign information
// @Description used to find campaign information
// @Tags briefs
// @Accept json
// @Produce json
// @Param id path string true "Brief ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.BriefInfo} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /briefs/info/{id} [get]
// @Security ApiKeyAuth
func GetBriefInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	info, err := r.GetBriefInfo(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, info)
}
