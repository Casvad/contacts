package controllers

import (
	"contacts/dto"
	"contacts/utils/constants"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getClaims(context *gin.Context) *dto.Claims {

	data, _ := context.Get(constants.Claims)

	return data.(*dto.Claims)
}

func getId[T int64 | int | int16 | int8](context *gin.Context) T {
	idStr := context.Param("id")

	res, _ := strconv.Atoi(idStr)

	return T(res)
}

func handleRequestWithBody[T interface{}, V interface{}](context *gin.Context, execution func(T) (V, error)) {
	var request T
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerResponse, err := execution(request)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, registerResponse)
}

func handleSimpleRequest[T interface{}](context *gin.Context, execution func() (T, error)) {

	registerResponse, err := execution()

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, registerResponse)
}
