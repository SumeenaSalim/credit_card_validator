package handler

import (
	"log"
	"net/http"

	"github.com/credit_card_validator/internal/model"
	"github.com/credit_card_validator/internal/service"
	"github.com/credit_card_validator/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/ggwhite/go-masker"
)


func CardHandler(c *gin.Context) {
	var card model.Card

	err := c.ShouldBindJSON(&card)
	if err != nil || card.Number == "" {
        log.Printf("[ERROR] Credit Card Validation - Invalid Request\n")
		utils.SendErrorResponse(c, http.StatusBadRequest, "'number' is required")
		return
	}

	number := masker.String(masker.MCreditCard, card.Number)

    if!service.CheckCard(card.Number) {
		log.Printf("[ERROR] Credit Card Validation - Card Number: '%s' is Invalid\n", number)
        utils.SendErrorResponse(c, http.StatusBadRequest, "'number' is invalid")
        return
    }

	log.Printf("[INFO] Credit Card Validation - Card Number: '%s' is Valid\n", number)
    utils.SendSuccessResponse(c, http.StatusOK, "'number' is valid")
}
