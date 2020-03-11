package access_token

import (
	"fmt"
	"github.com/mojoboss/bookstore_users-api/utils/crypto_utils"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"strings"
	"time"
)

const (
	ExpirationTime             = 24
	GrantTypePassword          = "password"
	GrantTypeClientCredentials = "client_credentials"
)

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	if at.GrantType != GrantTypePassword && at.GrantType != GrantTypeClientCredentials {
		return errors.NewBadRequestError("Invalid grant type")
	}
	//TODO: Validate each new grant type
	return nil
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	//used for password grant type
	Email    string `json:"email"`
	Password string `json:"password"`
	//used for client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(ExpirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMD5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
