package services

import "github.com/mrehanabbasi/company-inc/models"

func (s *Service) AddCompany(company *models.Company) (*models.Company, error) {
	return s.db.AddCompany(company)
}

func (s *Service) UpdateCompany(id string, companyUpdate *models.CompanyUpdate) (*models.Company, error) {
	return s.db.UpdateCompany(id, companyUpdate)
}

func (s *Service) GetCompanyByID(id string) (*models.Company, error) {
	return s.db.GetCompanyByID(id)
}

func (s *Service) DeleteCompany(id string) error {
	return s.db.DeleteCompany(id)
}
