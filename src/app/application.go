package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mojoboss/bookstore_oauth-api/src/clients/cassandra"
	"github.com/mojoboss/bookstore_oauth-api/src/http"
	"github.com/mojoboss/bookstore_oauth-api/src/repository/db"
	"github.com/mojoboss/bookstore_oauth-api/src/services"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	dbRepository := db.NewRepository()
	atService := services.NewService(dbRepository)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)
	router.Run(":8080")
}
