package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	Data    string `json:"data"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Response{message, "", ""})
}

func newSuccessResponse(method string, name string) {
	if name == "" {
		logrus.Printf("Succesful request for %s", method)
	} else {
		logrus.Printf("Succesful request for %s - %s", method, name)
	}
}
