package telegram

import (
	"net/url"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"encoding/base64"
)

type Client struct {
	token string
}

func (cl *Client) Call(method string, params url.Values, v interface{}) {
	apiUrl, _ := buildApiUrl(BaseURL, cl.Bot(), method)
	apiUrl.RawQuery = params.Encode()

	resp, _ := http.Get(apiUrl.String())
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(v)
}

func (cl *Client) GetMe() Account {
	resp := AccountResponse{}
	cl.Call("getMe", url.Values{}, &resp)
	return resp.Account
}

func (cl *Client) GetFile(fileId string) File  {
	resp := FileResponse{}
	cl.Call("getFile", url.Values{
		"file_id": []string{ fileId },
	}, &resp)

	return resp.File
}

func (cl *Client) ReadFile(fileId string) File {
	file := cl.GetFile(fileId)
	apiUrl, _ := buildApiUrl(BaseURL, "file", cl.Bot(), file.Path)
	resp, _ := http.Get(apiUrl.String())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	file.Content = "data:image/png;base64," + base64.StdEncoding.EncodeToString(body)
	return file
}

func (cl *Client) Bot() string {
	return "bot" + cl.token
}