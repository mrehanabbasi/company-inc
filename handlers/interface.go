package handlers

import "github.com/mrehanabbasi/company-inc/services"

type Handler struct {
	Service services.Service
}

func NewHandler(service services.Service) Handler {
	return Handler{
		Service: service,
	}
}
