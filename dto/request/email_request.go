package dto

import (
	"encoding/json"
	"time"
)

type EmailRequest struct {
	From    string    `json:"from" validate:"required"`
	To      []string  `json:"to" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
	Subject string    `json:"subject" validate:"required"`
	Content string    `json:"content" validate:"required"`
}

func ValidateEmailRequest(req *EmailRequest, body []byte) error {
	if err := json.Unmarshal(body, &req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}
