package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}


func SendSuccessResponse(c *gin.Context, status int, result string) {
	c.JSON(status, SuccessResponse{
		Status: "success",
		Result: result,
	})
}

func SendErrorResponse(c *gin.Context, status int, err string) {
	c.JSON(status, ErrorResponse{
		Status: "error",
		Error:  err,
	})
}
