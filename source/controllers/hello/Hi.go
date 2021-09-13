package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hi(context *gin.Context) {
	context.String(http.StatusOK, "hello there again")
}
