package database

import (
	"context"

	"github.com/pkg/errors"

	"github.com/gofrs/uuid"
	domainErr "github.com/mrehanabbasi/company-inc/errors"
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

func (m Client) UpdateCompany(id string, companyUpdate *models.CompanyUpdate) (*models.Company, error) {
	collection := m.GetMongoCompanyCollection()

	// Fetch the existing company from the DB
	company, err := m.GetCompanyByID(id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Apply the updates
	if companyUpdate.Name != nil {
		company.Name = *companyUpdate.Name
	}
	if companyUpdate.Description != nil {
		company.Description = *companyUpdate.Description
	}
	if companyUpdate.EmpCount != nil {
		company.EmpCount = *companyUpdate.EmpCount
	}
	if companyUpdate.IsRegistered != nil {
		company.IsRegistered = *companyUpdate.IsRegistered
	}
	if companyUpdate.Type != nil {
		company.Type = *companyUpdate.Type
	}

	// Updating the record in the DB
	company.ID = id

	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": company.ID}}, bson.M{"$set": company}); err != nil {
		log.Error(err.Error())
		return nil, errors.Wrap(err, "failed to updated company record")
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

func (m Client) DeleteCompany(id string) error {
	collection := m.GetMongoCompanyCollection()
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Error(err.Error())
		return errors.Wrap(err, "failed to delete company")
	}
	if res.DeletedCount == 0 {
		err = domainErr.NewAPIError(domainErr.NotFound, "company with the given id not found")
		log.Error(err.Error())
		return err
	}

	return nil
}
