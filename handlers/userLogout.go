package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrehanabbasi/company-inc/config"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/spf13/viper"
)

func (h *Handler) UserLogout(c *gin.Context) {
	cookieName := viper.GetString(config.CookieName)
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	// Check if the cookie was actually deleted
	_, err := c.Request.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		log.Error("failed to log out")
		c.JSON(http.StatusInternalServerError, domainErr.NewAPIError(domainErr.InternalServerError, "failed to log out"))
		return
	}

	c.Status(http.StatusNoContent)
}
