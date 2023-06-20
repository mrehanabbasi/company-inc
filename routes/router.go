package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrehanabbasi/company-inc/database"
	"github.com/mrehanabbasi/company-inc/handlers"
	"github.com/mrehanabbasi/company-inc/msq"
	"github.com/mrehanabbasi/company-inc/services"
)

func NewRouter(dbClient *database.Client, kConn *msq.MsqConn) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	handler := handlers.NewHandler(*services.NewService(dbClient), kConn)

	v1 := router.Group("/v1")
	{
		v1.POST("/signup", handler.UserSignup)
		v1.POST("/login", handler.UserLogin)
		v1.POST("/logout", handler.CheckAuth, handler.UserLogout)
	}

	companies := v1.Group("/companies")
	{
		companies.POST("", handler.CheckAuth, handler.AddCompany)
		companies.GET("/:id", handler.GetCompanyByID)
		companies.PATCH("/:id", handler.CheckAuth, handler.UpdateCompany)
		companies.DELETE("/:id", handler.CheckAuth, handler.DeleteCompany)
	}

	return router
}
