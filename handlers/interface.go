package handlers

import (
	"github.com/mrehanabbasi/company-inc/msq"
	"github.com/mrehanabbasi/company-inc/services"
)

type Handler struct {
	Service services.Service
	MsqConn *msq.MsqConn
}

func NewHandler(service services.Service, kafkaConn *msq.MsqConn) Handler {
	return Handler{
		Service: service,
		MsqConn: kafkaConn,
	}
}
