package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrehanabbasi/company-inc/constants"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
)

func (h *Handler) AddCompany(c *gin.Context) {
	var company *models.Company
	err := c.ShouldBindJSON(company)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, constants.NewAPIError(constants.BadRequest, err.Error()))
		return
	}

	company, err = h.CompanyService.AddCompany(company)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, constants.NewAPIError(constants.InternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, company)
}
