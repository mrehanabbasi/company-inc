package services

import "github.com/mrehanabbasi/company-inc/models"

func (s *Service) AddCompany(company *models.Company) (*models.Company, error) {
	return s.db.AddCompany(company)
}
