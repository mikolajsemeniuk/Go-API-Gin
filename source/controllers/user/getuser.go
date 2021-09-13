package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "get user")
}
