package prometheus

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Query(baseURL string, query string, start string, end string, step string) (PromResponse, error) {
	var response PromResponse

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
	}

	u, err := url.Parse(baseURL + "/api/v1/query_range")
	if err != nil {
		return response, err
	}

	params := u.Query()
	params.Set("query", query)
	params.Set("start", start)
	params.Set("end", end)
	params.Set("step", step)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return response, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, err
	}

	if response.Status != "success" {
		return response, fmt.Errorf("Prometheus query failed with status: %s", response.Status)
	}

	return response, nil
}
