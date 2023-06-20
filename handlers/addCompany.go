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
	err := c.ShouldBindJSON(&company)
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

	company, err = h.Service.AddCompany(company)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domainErr.NewAPIError(domainErr.InternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, company)
}
