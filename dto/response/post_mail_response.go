package dto

type PostMailResponse struct {
	Message     string `json:"message"`
	TopLevelID  string `json:"id"`
	ID          string `json:"_id"`
	Index       string `json:"_index"`
	Version     int64  `json:"_version"`
	SeqNo       int64  `json:"_seq_no"`
	PrimaryTerm int64  `json:"_primary_term"`
	Result      string `json:"result"`
}
