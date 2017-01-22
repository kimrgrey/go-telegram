package telegram

import (
	"strings"
	"net/url"
)

const BaseURL = "https://api.telegram.org"

func NewClient(token string) *Client {
	return &Client{token}
}

func buildApiUrl(parts ...string) (*url.URL, error) {
	rawurl := strings.Join(parts, "/")
	return url.Parse(rawurl)
}
