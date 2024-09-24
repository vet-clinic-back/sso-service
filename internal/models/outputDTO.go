package models

type SuccessDTO struct {
	Token string `json:"token"`
}

type ErrorDTO struct {
	Message string `json:"message"`
}
