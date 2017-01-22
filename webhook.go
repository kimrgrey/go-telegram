package telegram

import (
	"io"
	"encoding/json"
)

type Webhook struct {
	ID int `json:"update_id"`
	Message Message `json:"message"`
}

func (cl *Client) ParseWebhook(r io.Reader) (Webhook, error) {
	wh := Webhook{}
	decoder := json.NewDecoder(r)
	error := decoder.Decode(&wh)
	return wh, error
}


func (wh Webhook) BestPhotoVersion() PhotoVersion {
	var best PhotoVersion

	for _, ver := range wh.Message.Photo {
		if ver.FileSize >= best.FileSize {
			best = ver
		}
	}

	return best
}
