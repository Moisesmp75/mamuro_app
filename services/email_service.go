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
	search_url = "http://localhost:4080/es/enron_mail/_search"
	post_url   = "http://localhost:4080/api/enron_mail/_doc"
	username   = "admin"
	password   = "Complexpass#123"
)

func errors(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	response := resp.ErrorResponse(message)
	w.Write(ToJSONresponse(response))
}

func ToJSON(value any) ([]byte, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func ToJSONresponse(response any) []byte {
	jsonData, err := json.Marshal(response)
	if err != nil {
		return []byte("Error to convert json")
	}
	return jsonData
}

func performHTTPPostRequest(url string, request string) ([]byte, error) {
	reqs, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
	if err != nil {
		return nil, err
	}
	reqs.SetBasicAuth(username, password)
	reqs.Header.Set("Content-Type", "application/json")

	resps, err := http.DefaultClient.Do(reqs)
	if err != nil {
		return nil, err
	}
	defer resps.Body.Close()

	bodyResponse, err := io.ReadAll(resps.Body)
	if err != nil {
		return nil, err
	}
	return bodyResponse, nil
}

func parseJSONtoDto(body []byte, response any) error {
	err := json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}

func buildSearchResponse(response models.SearchDataResponse, request req.SearchRequest) resp.BaseResponse[[]models.Source] {
	total_items := response.Hits.Total.Value
	size := request.Size
	from := request.From
	meta_pagination := common.GenerateMeta(total_items, size, from)
	data_result := response.Hits.Hits
	mails := models.MapResource(data_result, models.HitsToSource)
	response_mail := resp.NewResponsePagination[[]models.Source](mails, meta_pagination)
	return response_mail
}

func SearchData(w http.ResponseWriter, r *http.Request) {
	bodyRequest := make([]byte, r.ContentLength)
	r.Body.Read(bodyRequest)

	var request req.SearchRequest
	if err := req.ValidateRequest(&request, bodyRequest); err != nil {
		errors(w, 400, err.Error())
		return
	}

	bodyResponse, err := performHTTPPostRequest(search_url, req.RequestToString(request))
	if err != nil {
		errors(w, 500, err.Error())
		return
	}

	var response models.SearchDataResponse
	if err := parseJSONtoDto(bodyResponse, &response); err != nil {
		errors(w, 500, err.Error())
		return
	}

	response_mail := buildSearchResponse(response, request)
	w.Write(ToJSONresponse(response_mail))
}

func PostMail(w http.ResponseWriter, r *http.Request) {
	bodyRequest := make([]byte, r.ContentLength)
	r.Body.Read(bodyRequest)

	var request req.EmailRequest
	if err := req.ValidateEmailRequest(&request, bodyRequest); err != nil {
		errors(w, 400, err.Error())
		return
	}

	reqs, err := ToJSON(request)
	if err != nil {
		errors(w, 400, err.Error())
		return
	}
	bodyResponse, err := performHTTPPostRequest(post_url, string(reqs))
	if err != nil {
		errors(w, 500, err.Error())
		return
	}

	var response resp.PostMailResponse
	if err := parseJSONtoDto(bodyResponse, &response); err != nil {
		errors(w, 500, err.Error())
		return
	}

	response_post := resp.NewResponse[resp.PostMailResponse](response)

	w.Write(ToJSONresponse(response_post))
}
