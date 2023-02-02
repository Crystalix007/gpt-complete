package complete

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Complete(input string) (string, error) {
	b := Body{
		Model:       "text-davinci-003",
		Prompt:      input,
		Temperature: 0,
		MaxTokens:   256,
	}
	j, err := b.JSON()
	if err != nil {
		return "", err
	}

	r, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", strings.NewReader(j))
	if err != nil {
		return "", err
	}

	p := Headers{
		ContentType: "application/json",
		APIKey:      API_KEY,
	}
	p.AddHeaders(&r.Header)

	c := http.DefaultClient

	res, err := c.Do(r)
	if err != nil {
		return "", err
	}

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resp Response

	if err := json.Unmarshal(bs, &resp); err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}
