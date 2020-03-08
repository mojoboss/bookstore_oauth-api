package services

import (
	"github.com/mojoboss/bookstore_oauth-api/src/domain/access_token"
	"github.com/mojoboss/bookstore_oauth-api/src/repository/db"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
}

type service struct {
	repository db.DbRepository
}

func NewService(repo db.DbRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessToken) *errors.RestErr {
	if err := request.Validate(); err != nil {
		return err
	}
	return s.repository.Create(request)
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
