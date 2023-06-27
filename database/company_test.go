package database

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/mrehanabbasi/company-inc/models"
)

func initTestSetup(ctx context.Context) *Client {
	mongoClient := InitTestDB(ctx)
	mongoClient.InitIndices()
	return mongoClient
}

func cleanUpTestEnv(client *Client) {
	_ = client.DeleteAllTestCompanies()
}

func TestAddCompany(t *testing.T) {
	// Setting up environment for testing
	ctx := context.TODO()
	client := initTestSetup(ctx)
	defer func() {
		_ = client.Conn.Disconnect(ctx)
	}()

	id, _ := uuid.NewV4()
	company := &models.Company{
		ID:           id.String(),
		Name:         "New Company 1",
		Description:  "Description of new company.",
		EmpCount:     100,
		IsRegistered: true,
		Type:         "Corporations",
	}

	company, err := client.AddCompany(company)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if company.ID != id.String() {
		t.Error("mongodb added new company id")
		return
	}

	id, _ = uuid.NewV4()
	company.ID = id.String()
	company, err = client.AddCompany(company)
	if err == nil {
		t.Error("duplicate company name should give error")
		return
	}

	// Cleaning up the environment
	cleanUpTestEnv(client)
}

func TestGetCompany(t *testing.T) {
	// Setting up environment for testing
	ctx := context.TODO()
	client := initTestSetup(ctx)
	defer func() {
		_ = client.Conn.Disconnect(ctx)
	}()

	id, _ := uuid.NewV4()
	company := &models.Company{
		ID:           id.String(),
		Name:         "New Company 1",
		Description:  "Description of new company.",
		EmpCount:     100,
		IsRegistered: true,
		Type:         "Corporations",
	}
	company, err := client.AddCompany(company)
	if err != nil {
		t.Error(err.Error())
		return
	}

	newCompany, err := client.GetCompanyByID(id.String())
	if newCompany != company {
		t.Error("could not get the correct company via id")
		return
	}

	// Cleaning up the environment
	cleanUpTestEnv(client)
}
