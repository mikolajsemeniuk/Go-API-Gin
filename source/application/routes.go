package application

import "Go-API/source/controllers"

func Route() {
	router.GET("/hello", controllers.Hello)
}
