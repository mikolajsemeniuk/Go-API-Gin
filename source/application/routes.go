package application

import "Go-API-Gin/source/controllers"

func Route() {
	router.GET("/hello", controllers.Hello)
}
