package app

import (
	"github.com/alessandroarosio/bookstore_oauth-api/src/domain/access_token"
	"github.com/alessandroarosio/bookstore_oauth-api/src/http"
	"github.com/alessandroarosio/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	repo := db.NewRepository()
	atService := access_token.NewService(repo)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
