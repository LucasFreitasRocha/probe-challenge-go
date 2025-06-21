package controller

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/config/rest_err"
	"github.com/gin-gonic/gin"
)

func notValidPayload(c *gin.Context) {
	errorMessage := rest_err.NewBadRequestError(
		"payload is not valid ",
	)
	c.JSON(errorMessage.Code, errorMessage)
}