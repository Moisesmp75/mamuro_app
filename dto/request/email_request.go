package dto

import (
	"encoding/json"
	"time"
)

type EmailRequest struct {
	From    string    `validate:"required"`
	To      []string  `validate:"required"`
	Date    time.Time `validate:"required"`
	Subject string    `validate:"required"`
	Content string    `validate:"required"`
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
