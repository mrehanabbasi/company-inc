package database

import (
	"context"

	"github.com/gofrs/uuid"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (m Client) UserSignup(user *models.User) (*models.User, error) {
	newUUID, _ := uuid.NewV4()
	user.ID = newUUID.String()

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return nil, errors.Wrap(err, "failed to hash password")
	}
	user.Password = string(hashedPassword)

	collection := m.GetMongoUserCollection()

	if _, err := collection.InsertOne(context.TODO(), user); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return user, nil
}
