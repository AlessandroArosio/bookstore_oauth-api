package rest

import (
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@email.com","password":"pass"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := NewRestUsersRepository()

	user, err := repository.LoginUser("email@email.com", "pass")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	// todo
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	// todo
}

func TestLoginInvalidUserJsonResponse(t *testing.T) {
	// todo
}

func TestLoginUserNoError(t *testing.T) {
	// todo
}
