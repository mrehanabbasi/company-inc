package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
)

func (h *Handler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteCompany(id); err != nil {
		log.Error(err.Error())
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			c.JSON(http.StatusNotFound, apiErr)
		default:
			c.JSON(http.StatusInternalServerError, domainErr.NewAPIError(domainErr.InternalServerError, err.Error()))
		}
		return
	}

	// Product message queue event
	if err := h.MsqConn.ProduceCompanyEvent(&models.Company{ID: id}, http.MethodDelete); err != nil {
		log.Error(err.Error())
	}

	c.Status(http.StatusNoContent)
}
