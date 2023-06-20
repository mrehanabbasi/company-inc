package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrehanabbasi/company-inc/config"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"github.com/spf13/viper"
)

func (h *Handler) CheckAuth(c *gin.Context) {
	// Get the JWT token from the cookie
	cookie, err := c.Request.Cookie(viper.GetString(config.CookieName))
	if err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, domainErr.NewAPIError(domainErr.Unauthorized, "user not logged in"))
		return
	}
	tokenString := cookie.Value

	secretKey := viper.GetString(config.SecretKey)

	// Parsing the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, domainErr.NewAPIError(domainErr.Unauthorized, "unauthorized token"))
		return
	}

	// Checking token validity
	if !token.Valid {
		log.Error("invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, domainErr.NewAPIError(domainErr.Unauthorized, "invalid token"))
		return
	}

	// Getting claims
	claims, ok := token.Claims.(*models.JwtClaims)
	if !ok {
		log.Error("invalid jwt claims")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid jwt claims"})
		return
	}

	userInfo := models.TokenInfo{
		UserID:    claims.UserID,
		UserName:  claims.UserName,
		UserEmail: claims.UserEmail,
	}
	c.Set("userInfo", userInfo)
}
