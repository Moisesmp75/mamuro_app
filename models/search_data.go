package models

import (
	"encoding/json"
	"time"
)

type SearchDataResponse struct {
	Took     int64  `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Index     Index   `json:"_index"`
	Type      Type    `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	Timestamp string  `json:"@timestamp"`
	Source    Source  `json:"_source"`
}

type Source struct {
	Timestamp               string                  `json:"@timestamp"`
	Bcc                     []string                `json:"bcc"`
	Cc                      []string                `json:"cc"`
	Content                 string                  `json:"content"`
	ContentTransferEncoding ContentTransferEncoding `json:"content_transfer_encoding"`
	ContentType             ContentType             `json:"content_type"`
	Date                    time.Time               `json:"date"`
	From                    string                  `json:"from"`
	MessageID               string                  `json:"message_id"`
	MIMEVersion             MIMEVersion             `json:"mime_version"`
	Subject                 string                  `json:"subject"`
	To                      []string                `json:"to"`
	XBcc                    interface{}             `json:"x_bcc"`
	XCc                     interface{}             `json:"x_cc"`
	XFilename               string                  `json:"x_filename"`
	XFolder                 string                  `json:"x_folder"`
	XFrom                   string                  `json:"x_from"`
	XOrigin                 string                  `json:"x_origin"`
	XTo                     []string                `json:"x_to"`
}

type Total struct {
	Value int64 `json:"value"`
}

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}

type Index string

const (
	EnronMail Index = "enron_mail"
)

type ContentTransferEncoding string

const (
	QuotedPrintable ContentTransferEncoding = "quoted-printable\u000d"
	The7Bit         ContentTransferEncoding = "7bit\u000d"
)

type ContentType string

const (
	TextPlainCharsetANSIX341968 ContentType = "text/plain; charset=ANSI_X3.4-1968\u000d"
	TextPlainCharsetUsASCII     ContentType = "text/plain; charset=us-ascii\u000d"
)

type MIMEVersion string

const (
	The10 MIMEVersion = "1.0\u000d"
)

type Type string

const (
	Doc Type = "_doc"
)

func ConvertToResponse(body []byte, response *SearchDataResponse) string {
	err := json.Unmarshal(body, &response)
	if err != nil {
		return err.Error()
	}
	return ""
}

func HitsToSource(hit Hit) Source {
	return hit.Source
}

func MapResource(array []Hit, f func(Hit) Source) []Source {
	result := make([]Source, len(array))
	for i, v := range array {
		result[i] = f(v)
	}
	return result
}
