package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrehanabbasi/company-inc/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	handler := handlers.NewHandler()

	v1 := router.Group("/v1")

	companies := v1.Group("/companies")
	{
		companies.POST("", handler.AddCompany)
		companies.GET("/:id", handler.GetCompanyByID)
		companies.PATCH("/:id", handler.UpdateCompany)
		companies.DELETE("/:id", handler.DeleteCompany)
	}

	return router
}
