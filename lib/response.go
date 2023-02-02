package complete

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}
