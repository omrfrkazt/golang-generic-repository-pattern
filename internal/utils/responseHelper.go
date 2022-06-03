package utils

import (
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/models"

	"github.com/gofiber/fiber/v2"
)

//BaseResponseModel helpers

//custom response
func NewResponse(status int, data interface{}, message string, success bool) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  status,
		Data:    data,
		Message: message,
		Success: success,
	}
}

//returns http 200 OK
func StatusOK(data interface{}) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusOK,
		Data:    data,
		Message: "OK",
		Success: true,
	}
}

//returns http 400
func StatusFail(message string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusBadRequest,
		Message: message,
		Success: false,
	}
}

// retuns http 401
func StatusUnauthorized(message string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusUnauthorized,
		Message: message,
		Success: false,
	}
}

//returns http 500
func UnhandledError() models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusInternalServerError,
		Message: "Unhandled error occurred. Please try again later",
		Success: false,
	}
}

//returns http 404
func StatusNotFound(message string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusNotFound,
		Message: message,
		Success: false,
	}
}
