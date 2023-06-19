package database

import (
	"context"

	"github.com/pkg/errors"

	"github.com/gofrs/uuid"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m Client) GetCompanyByID(id string) (*models.Company, error) {
	var company *models.Company
	collection := m.GetMongoCompanyCollection()

	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&company); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Error(err.Error())
			return nil, errors.Wrap(err, "failed to fetch company... not found")
		}
		return nil, err
	}

	return company, nil
}
