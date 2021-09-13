package application

import (
	"Go-API-Gin/source/controllers/hello"
	"Go-API-Gin/source/controllers/user"
)

func Route() *Application {
	router.GET("/hello", hello.Hi)
	router.GET("/user", user.GetUser)
	router.POST("/user", user.CreateUser)
	return &Application{}
}
