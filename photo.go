package telegram

type Photo []PhotoVersion

type PhotoVersion struct {
	FileID string `json:"file_id"`
	FileSize int `json:"file_size"`
	Width int `json:"width"`
	Height int `json:"height"`
}