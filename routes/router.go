package routes

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	v1 := router.Group("/v1")

	companies := v1.Group("/companies")
	{
		companies.POST("")
	}

	return router
}
