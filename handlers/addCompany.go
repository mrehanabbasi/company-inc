package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
)

func (h *Handler) AddCompany(c *gin.Context) {
	var company *models.Company
	err := c.ShouldBindJSON(company)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, domainErr.NewAPIError(domainErr.BadRequest, err.Error()))
		return
	}

	company, err = h.CompanyService.AddCompany(company)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domainErr.NewAPIError(domainErr.InternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, company)
}
