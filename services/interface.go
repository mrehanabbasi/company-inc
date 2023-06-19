package services

import (
	"github.com/mrehanabbasi/company-inc/database"
)

type Service struct {
	db *database.Client
}

func NewService(client *database.Client) *Service {
	return &Service{db: client}
}
