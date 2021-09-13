package application

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func (*Application) Listen(port ...string) {
	if len(port) == 0 {
		router.Run()
	} else {
		router.Run(port[0])
	}
}
