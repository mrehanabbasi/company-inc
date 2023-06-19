package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
)

func (h *Handler) UpdateCompany(c *gin.Context) {
	id := c.Param("id")

	var companyUpdate *models.CompanyUpdate
	err := c.ShouldBindJSON(companyUpdate)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, domainErr.NewAPIError(domainErr.BadRequest, err.Error()))
		return
	}

	company, err := h.CompanyService.UpdateCompany(id, companyUpdate)
	if err != nil {
		log.Error(err.Error())
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			c.JSON(http.StatusNotFound, domainErr.NewAPIError(domainErr.NotFound, err.Error()))
			return
		default:
			c.JSON(http.StatusInternalServerError, domainErr.NewAPIError(domainErr.InternalServerError, err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, company)
}
