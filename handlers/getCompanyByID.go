package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
)

func (h *Handler) GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	company, err := h.Service.GetCompanyByID(id)
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
