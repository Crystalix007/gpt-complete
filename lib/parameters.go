package complete

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Headers struct {
	ContentType string
	APIKey      string
}

func (o Headers) AddHeaders(h *http.Header) {
	h.Add("Content-Type", o.ContentType)
	h.Add("Authorization", fmt.Sprintf("Bearer %s", o.APIKey))
}

type Body struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	Temperature uint   `json:"temperature"`
	MaxTokens   uint   `json:"max_tokens"`
}

func (b Body) JSON() (string, error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
