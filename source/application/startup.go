package application

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartUp() {
	Route()
	router.Run()
}
