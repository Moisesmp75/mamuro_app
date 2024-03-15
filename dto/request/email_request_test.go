package dto

import (
	"fmt"
	"testing"
	"time"
)

func compareEmailRequests(req1, req2 *EmailRequest) bool {
	if req1.From != req2.From {
		return false
	}
	if req1.Date != req2.Date {
		return false
	}
	if len(req1.To) != len(req2.To) {
		return false
	}
	for i := range req1.To {
		if req1.To[i] != req2.To[i] {
			return false
		}
	}
	if req1.Subject != req2.Subject {
		return false
	}
	if req1.Content != req2.Content {
		return false
	}
	return true
}

func TestValidateEmailRequest(t *testing.T) {
	// Case 1: Valid JSON and all fields present
	validJSON1 := `{"from":"example@example.com","to":["recipient@example.com"],"date":"2024-03-20T12:00:00-08:00","subject":"Test","content":"This is a test email"}`
	req1 := &EmailRequest{}
	if err := ValidateEmailRequest(req1, []byte(validJSON1)); err != nil {
		t.Errorf("Unexpected error for valid JSON: %v", err)
	}

	// Case 2: Valid JSON bu missing required filed
	invalidJSON2 := `{"to":["recipient@example.com"],"date":"2024-03-20T12:00:00Z","subject":"Test","content":"This is a test email"}`
	req2 := &EmailRequest{}
	if err := ValidateEmailRequest(req2, []byte(invalidJSON2)); err == nil {
		fmt.Println(err)
		t.Error("Expected error for missing 'from' field but got nil")
	}

	// Case 3: Invalid JSON
	invalidJSON3 := `{}`
	req3 := &EmailRequest{}
	if err := ValidateEmailRequest(req3, []byte(invalidJSON3)); err == nil {
		t.Error("Expected error for invalid JSON but got nil")
	}

	// Case 4: Valid JSON but invalid data
	invalidJSON4 := `{"from":"example@example.com","to":["recipient@example.com"],"date":"invalid date","subject":"Test","content":"This is a test email"}`
	req4 := &EmailRequest{}
	if err := ValidateEmailRequest(req4, []byte(invalidJSON4)); err == nil {
		t.Error("Expected error for invalid data but got nil")
	}
	// Case 5: Valid JSON but invalid fields
	invalidJSON5 := `{"from":"example@example.com","to":["example@example.com", "another@example.com"],"date":"2024-03-20T12:00:00-08:00","subject":"Test","content":"This is a test email", "cc": "example@example.com", "MessageID": "id122222"}`
	req5 := &EmailRequest{}
	date, _ := time.Parse(time.RFC3339, "2024-03-20T12:00:00-08:00")
	expectedReq5 := EmailRequest{
		From:    "example@example.com",
		Date:    date,
		To:      []string{"example@example.com", "another@example.com"},
		Subject: "Test",
		Content: "This is a test email",
	}
	err := ValidateEmailRequest(req5, []byte(invalidJSON5))
	if err != nil {
		t.Error("Expected error for invalid data but got nil")
	}
	if !compareEmailRequests(req5, &expectedReq5) {
		t.Errorf("Structure req5 does not match expectedReq5. Expected: %v but got: %v", expectedReq5, *req5)
	}

}
