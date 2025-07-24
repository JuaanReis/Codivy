package requests

import (
	"io"
	"fmt"
	"net/http"
	"Codivy/pkg"
)

func (c *Client) Post(endpoint string, body any) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	jsonData, err := pkg.ToJSON(body)
	if err != nil {
		return nil, fmt.Errorf("error serializing JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}

	return resp, nil
}