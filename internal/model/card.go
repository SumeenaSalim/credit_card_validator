package model

type Card struct {
	Number string `json:"number" validate:"required"`
}
