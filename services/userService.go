package services

import "github.com/mrehanabbasi/company-inc/models"

func (s *Service) UserSignup(user *models.User) (*models.User, error) {
	return s.db.UserSignup(user)
}

func (s *Service) UserLogin(loginInfo *models.UserLogin) (string, error) {
	return s.db.UserLogin(loginInfo)
}
