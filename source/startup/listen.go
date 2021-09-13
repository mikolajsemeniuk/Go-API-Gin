package startup

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func (startup *Startup) Listen(port ...string) {
	if len(port) == 0 {
		router.Run()
	} else {
		router.Run(port[0])
	}
}
