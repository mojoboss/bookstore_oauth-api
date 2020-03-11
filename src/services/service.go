package services

import (
	"github.com/mojoboss/bookstore_oauth-api/src/domain/access_token"
	"github.com/mojoboss/bookstore_oauth-api/src/repository/db"
	"github.com/mojoboss/bookstore_oauth-api/src/repository/rest"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
}

type service struct {
	repository     db.DbRepository
	userRepository rest.RestUsersRepository
}

func NewService(repo db.DbRepository, userRepository rest.RestUsersRepository) Service {
	return &service{
		repository:     repo,
		userRepository: userRepository,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	if request.GrantType == access_token.GrantTypePassword {
		user, err := s.userRepository.LoginUser(request.Email, request.Password)
		if err != nil {
			return nil, err
		}
		at := access_token.GetNewAccessToken(user.Id)
		at.Generate()
		if err = s.repository.Create(at); err != nil {
			return nil, err
		}
	}
	return nil, errors.NewBadRequestError("GrantType not supported")
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
