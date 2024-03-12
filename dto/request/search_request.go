package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SearchRequest struct {
	Query string
	Size  int   `validate:"gte=0"`
	From  int64 `validate:"gte=0"`
	Sort  string
}

var (
	validate = validator.New()
)

func ValidateRequest(req SearchRequest) string {
	if err := validate.Struct(req); err != nil {
		return err.Error()
	}
	return ""
}

func RequestToString(req SearchRequest) string {
	if req.Sort == "" {
		req.Sort = "-date"
	}
	if req.Size == 0 {
		req.Size = 10
	}
	if req.From == 0 {
		req.From = 0
	}
	if req.Query == "" {
		return parseRequestGetAll(req)
	}
	return parseRequestSearch(req)
}

func parseRequestGetAll(req SearchRequest) string {
	return fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{
						"match_all": {}
					}
				]
			}
		},
		"sort": [
			"-%v"
		],
		"from": %v,
		"size": %v
	}
	`, req.Sort, req.From, req.Size)
}

func parseRequestSearch(req SearchRequest) string {
	return fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{
						"query_string": {
							"query": "%v"
						}
					}
				]
			}
		},
		"sort": [
			"%v"
		],
		"from": %v,
		"size": %v
	}
	`, req.Query, req.Sort, req.From, req.Size)
}
