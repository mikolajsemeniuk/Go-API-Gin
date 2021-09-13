package startup

import (
	"Go-API-Gin/source/controllers/hello"
	"Go-API-Gin/source/controllers/user"
)

func Route() Startup {
	router.GET("/hello", hello.Hi)
	router.GET("/user", user.GetUser)
	router.POST("/user", user.CreateUser)
	return Startup{}
}
