package services

import (
	"bytes"
	"encoding/json"
	"io"
	"mamuro_app/common"
	req "mamuro_app/dto/request"
	resp "mamuro_app/dto/response"
	"mamuro_app/models"
	"net/http"
)

var (
	url      = "http://zincsearch:4080/es/enron_mail/_search"
	username = "admin"
	password = "Complexpass#123"
)

func errors(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	response := resp.ErrorResponse(message)
	w.Write([]byte(resp.JsonResponse[*string](response)))
}

func validateRequest(w http.ResponseWriter, request *req.SearchRequest, body []byte) bool {
	if err := json.Unmarshal(body, &request); err != nil {
		errors(w, 400, err.Error())
		return false
	}
	if err := req.ValidateRequest(*request); err != "" {
		errors(w, 400, err)
		return false
	}
	return true
}

func SearchData(w http.ResponseWriter, r *http.Request) {
	bodyRequest := make([]byte, r.ContentLength)
	r.Body.Read(bodyRequest)

	var request req.SearchRequest
	if !validateRequest(w, &request, bodyRequest) {
		return
	}

	reqs, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(req.RequestToString(request))))
	if err != nil {
		errors(w, 500, err.Error())
		return
	}
	reqs.SetBasicAuth(username, password)
	reqs.Header.Set("Content-Type", "application/json")

	resps, err := http.DefaultClient.Do(reqs)
	if err != nil {
		errors(w, 500, err.Error())
		return
	}
	defer resps.Body.Close()

	bodyResponse, err := io.ReadAll(resps.Body)
	if err != nil {
		errors(w, 500, err.Error())
		return
	}

	var response models.SearchDataResponse
	if err := models.ConvertToResponse(bodyResponse, &response); err != "" {
		errors(w, 500, err)
		return
	}

	total_items := response.Hits.Total.Value
	size := request.Size
	from := request.From
	meta_pagination := common.GenerateMeta(total_items, size, from)
	data_result := response.Hits.Hits
	mails := models.MapResource(data_result, models.HitsToSource)
	response_mail := resp.NewResponsePagination[[]models.Source](mails, meta_pagination)
	w.Write([]byte(resp.JsonResponse[[]byte](response_mail)))
}
