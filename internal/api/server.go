package api

import (
	"log"

	"github.com/credit_card_validator/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	r := gin.Default()

	r.POST("/validate", handler.CardHandler)
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server '8080': ", err)
	} else {
		log.Fatal("Server is listening on: 8080")
	}
}