package common

import "testing"

func TestGenerateMeta(t *testing.T) {
	//Case 1
	totalItems := int64(100)
	size := 10
	from := int64(0)
	expectedMeta := Meta{
		TotalItems:   100,
		ItemsPerPage: 10,
		CurrentPage:  1,
		TotalPages:   10,
		HasNextPage:  true,
		HasPrevPage:  false,
	}
	meta := GenerateMeta(totalItems, size, from)
	if meta != expectedMeta {
		t.Errorf("Expected %v but got %v", expectedMeta, meta)
	}
	//Case 2
	totalItems = 75
	size = 10
	from = 30
	expectedMeta = Meta{
		TotalItems:   75,
		ItemsPerPage: 10,
		CurrentPage:  4,
		TotalPages:   8,
		HasNextPage:  true,
		HasPrevPage:  true,
	}
	meta = GenerateMeta(totalItems, size, from)
	if meta != expectedMeta {
		t.Errorf("Expected %v but got %v", expectedMeta, meta)
	}
}
