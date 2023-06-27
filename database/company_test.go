package database

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/mrehanabbasi/company-inc/models"
)

func InitTestSetup(ctx context.Context) *Client {
	mongoClient := InitTestDB(ctx)
	mongoClient.InitIndices()
	return mongoClient
}

func TestAddCompany(t *testing.T) {
	ctx := context.TODO()
	client := InitTestSetup(ctx)
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
}
