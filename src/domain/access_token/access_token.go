package access_token

import (
	"fmt"
	"github.com/alessandroarosio/bookstore-utils-go/rest_errors"
	"github.com/alessandroarosio/bookstore_users-API/utils/crypto_utils"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atr *AccessTokenRequest) Validate() rest_errors.RestErr {
	switch atr.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant type")
	}
	return nil
}

func (at *AccessToken) Validate() rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	} else if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	} else if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	} else if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
