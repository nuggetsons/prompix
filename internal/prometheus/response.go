package prometheus

type PromResponse struct {
	Status string   `json:"status"`
	Data   PromData `json:"data"`
}

type PromData struct {
	ResultType string       `json:"resultType"`
	Result     []PromResult `json:"result"`
}

type PromResult struct {
	Metric map[string]string `json:"metric"`
	Values []any             `json:"values"`
}
