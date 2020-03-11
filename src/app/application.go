package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mojoboss/bookstore_oauth-api/src/http"
	"github.com/mojoboss/bookstore_oauth-api/src/repository/db"
	"github.com/mojoboss/bookstore_oauth-api/src/repository/rest"
	"github.com/mojoboss/bookstore_oauth-api/src/services"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	userRepo := rest.NewRestUsersRepository()
	atService := services.NewService(dbRepository, userRepo)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)
	router.Run(":8081")
}
