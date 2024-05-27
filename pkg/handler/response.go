package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, status int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(status, ErrorResponse{message})
}
