package telegram

type File struct {
	ID string `json:"file_id"`
	Size int `json:"file_size"`
	Path string `json:"file_path"`
	Content string `json:"-"`
}

type FileResponse struct {
	Ok bool `json:"ok"`
	File File `json:"result"`
}

