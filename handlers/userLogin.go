package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrehanabbasi/company-inc/config"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"github.com/spf13/viper"
)

func (h *Handler) UserLogin(c *gin.Context) {
	var login *models.UserLogin
	err := c.ShouldBindJSON(&login)
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

	token, err := h.Service.UserLogin(login)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusUnauthorized, domainErr.NewAPIError(domainErr.Unauthorized, err.Error()))
		return
	}

	// Creating cookie for login
	if err = setTokenCookie(c, token); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domainErr.NewAPIError(domainErr.InternalServerError, "unable to create cookie"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func setTokenCookie(c *gin.Context, token string) error {
	log.Info("Generating cookie from JWT token")

	tokenExpiry, err := time.ParseDuration(viper.GetString(config.TokenExpiry))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(tokenExpiry),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)
	return nil
}
