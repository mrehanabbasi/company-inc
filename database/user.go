package database

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrehanabbasi/company-inc/config"
	log "github.com/mrehanabbasi/company-inc/logger"
	"github.com/mrehanabbasi/company-inc/models"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m Client) UserLogin(loginInfo *models.UserLogin) (string, error) {
	user, err := m.GetUserFromEmail(loginInfo.Email)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	// Check if password is correct or not
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		log.Error(err.Error())
		return "", errors.Wrap(err, "incorrect password")
	}

	// Generate JWT Token
	token, err := generateToken(user.ID)
	if err != nil {
		log.Error(err.Error())
		return "", errors.Wrap(err, "failed to generate token")
	}

	return token, nil
}

func (m Client) GetUserFromEmail(email string) (*models.User, error) {
	var user *models.User

	collection := m.GetMongoUserCollection()

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Error(err.Error())
			return nil, errors.Wrap(err, "failed to fetch user... not found")
		}
		return nil, err
	}

	return user, nil
}

func generateToken(userID string) (string, error) {
	log.Info("Generating JWT token")

	tokenExpiry, err := time.ParseDuration(viper.GetString(config.TokenExpiry))
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	// Get the secret key from the env var
	secretKey := viper.GetString(config.SecretKey)

	//Set the token expiration time
	expirationTime := time.Now().Add(tokenExpiry)

	claims := &models.JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return token, nil
}
