package user

import (
	"net/http"

	"Go-API-Gin/source/entities"
	"Go-API-Gin/source/payloads"
	"Go-API-Gin/source/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var user entities.User
	// this is equal this below
	//
	// bytes, error := ioutil.ReadAll(context.Request.Body)
	// if error != nil {
	// 	return
	// }
	// if error := json.Unmarshal(bytes, &user); error != nil {
	// 	return
	// }
	if error := context.ShouldBindJSON(&user); error != nil {
		invalidBody := payloads.InvalidPayloadError{
			Message: "invalid json body",
			Code:    http.StatusBadRequest,
			Error:   "bad request",
		}
		context.JSON(invalidBody.Code, invalidBody)
		return
	}

	result, error := services.CreateUser(user)
	if error != nil {
		return
	}
	context.JSON(http.StatusCreated, result)
}
