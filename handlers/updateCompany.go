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
	err := c.ShouldBindJSON(&companyUpdate)
	if err != nil {
		log.Error(err.Error())
		errs, ok := models.ErrValidationSlice(err)
		if !ok {
			c.JSON(http.StatusBadRequest, domainErr.NewAPIError(domainErr.BadRequest, err.Error()))
			return
		}

		if len(errs) > 1 {
			c.JSON(http.StatusBadRequest, domainErr.NewAPIErrors(domainErr.BadRequest, errs))
		} else {
			c.JSON(http.StatusBadRequest, domainErr.NewAPIError(domainErr.BadRequest, errs[0]))
		}
		return
	}

	company, err := h.Service.UpdateCompany(id, companyUpdate)
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

	// Product message queue event
	if err = h.MsqConn.ProduceCompanyEvent(company, http.MethodPatch); err != nil {
		log.Error(err.Error())
	}

	c.JSON(http.StatusOK, company)
}
