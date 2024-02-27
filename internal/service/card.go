package service

import (
	"strings"

	"github.com/credit_card_validator/internal/utils"
)

func CheckCard(card string) bool {
	card = strings.Replace(card, " ", "", -1)

	if len(card) < 16 || len(card) > 19 {
		return false
	}

	return utils.LuhnCheck(card)
}