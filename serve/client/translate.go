package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// TranslateTranslatePayload is the data structure used to initialize the Translate translate request body.
type TranslateTranslatePayload struct {
	// Operand name
	Name string `json:"name" xml:"name"`
}

// Translate a Japanese word to English
func (c *Client) TranslateTranslate(path string, payload *TranslateTranslatePayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
