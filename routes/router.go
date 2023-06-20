package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrehanabbasi/company-inc/database"
	"github.com/mrehanabbasi/company-inc/handlers"
	"github.com/mrehanabbasi/company-inc/services"
)

func NewRouter(dbClient *database.Client) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	handler := handlers.NewHandler(*services.NewService(dbClient))

	v1 := router.Group("/v1")
	{
		v1.POST("/signup")
		v1.POST("/login")
		v1.POST("/logout")
	}

	companies := v1.Group("/companies")
	{
		companies.POST("", handler.AddCompany)
		companies.GET("/:id", handler.GetCompanyByID)
		companies.PATCH("/:id", handler.UpdateCompany)
		companies.DELETE("/:id", handler.DeleteCompany)
	}

	return router
}
