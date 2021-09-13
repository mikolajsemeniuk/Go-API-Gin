package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Not Implemented")
}
