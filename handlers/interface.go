package handlers

import "github.com/mrehanabbasi/company-inc/services"

type Handler struct {
	CompanyService services.Service
}

func NewHandler(companyService services.Service) Handler {
	return Handler{
		CompanyService: companyService,
	}
}
