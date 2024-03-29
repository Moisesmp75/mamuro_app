package dto

import (
	"fmt"
	"testing"
)

func compareSearchRequests(req1, req2 *SearchRequest) bool {
	if req1.Query != req2.Query {
		return false
	}
	if req1.Size != req2.Size {
		return false
	}
	if req1.From != req2.From {
		return false
	}
	if req1.Sort != req2.Sort {
		return false
	}
	return true
}

func TestValidateRequest(t *testing.T) {
	// Case 1: Valid JSON and all fields present
	validJSON1 := `{"query":"test","size":10,"from":0,"sort":"-date"}`
	req1 := &SearchRequest{}
	if err := ValidateRequest(req1, []byte(validJSON1)); err != nil {
		t.Errorf("Error inesperado para JSON válido: %v", err)
	}

	// Caso 2: Valid JSON but missing field
	invalidJSON2 := `{"query":"test","from":0,"sort":"-date"}`
	req2 := &SearchRequest{}
	expectedReq2 := SearchRequest{
		Query: "test",
		From:  0,
		Size:  0,
		Sort:  "-date",
	}
	err2 := ValidateRequest(req2, []byte(invalidJSON2))
	if err2 != nil {
		fmt.Println(req2)
		t.Error("Se esperaba un error para el campo faltante 'size' pero se obtuvo nil")
	}
	if !compareSearchRequests(req2, &expectedReq2) {
		t.Errorf("Structure req5 does not match expectedReq5. Expected: %v but got: %v", expectedReq2, *req2)
	}

	// Caso 3: Invalid JSON
	invalidJSON3 := `{}`
	req3 := &SearchRequest{}
	expectedReq3 := SearchRequest{
		From: 0,
		Size: 0,
	}
	err3 := ValidateRequest(req3, []byte(invalidJSON3))
	if err3 != nil {
		fmt.Println(req3)
		t.Error("Se esperaba un error para JSON inválido pero se obtuvo nil")
	}
	if !compareSearchRequests(req3, &expectedReq3) {
		t.Errorf("Structure req5 does not match expectedReq5. Expected: %v but got: %v", expectedReq2, *req2)
	}
}

func TestRequestToString(t *testing.T) {
	// Case 1: Request with empty Query
	req1 := SearchRequest{
		Query: "",
		Size:  5,
		From:  0,
		Sort:  "-date",
	}
	expected1 := parseRequestGetAll(req1)
	result1 := RequestToString(req1)
	if result1 != expected1 {
		t.Errorf("Case 1: Unexpected result. Expected: %s, but got: %s", expected1, result1)
	}

	// Case 2: Request with non-empty Query
	req2 := SearchRequest{
		Query: "test",
		Size:  10,
		From:  0,
		Sort:  "-date",
	}
	expected2 := parseRequestSearch(req2)
	result2 := RequestToString(req2)
	if result2 != expected2 {
		t.Errorf("Case 2: Unexpected result. Expected: %s, but got: %s", expected2, result2)
	}

	// Case 3: Request with empty Sort
	req3 := SearchRequest{
		Query: "test",
		Size:  10,
		From:  0,
		Sort: "-@timestamp",
	}
	expected3 := parseRequestSearch(req3)
	result3 := RequestToString(req3)
	if result3 != expected3 {
		t.Errorf("Case 3: Unexpected result. Expected: %s, but got: %s", expected3, result3)
	}
}