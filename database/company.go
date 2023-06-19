package database

import (
	"context"

	"github.com/gofrs/uuid"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
)

func (m Client) AddCompany(company *models.Company) (*models.Company, error) {
	newUUID, _ := uuid.NewV4()
	company.ID = newUUID.String()

	collection := m.GetMongoCompanyCollection()

	if _, err := collection.InsertOne(context.TODO(), company); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return company, nil
}
