package client

import (
	"context"
	"net/http"
	"time"
)

type HTTPDoctorClient struct {
	baseURL string
	client  *http.Client
}

func NewHTTPDoctorClient(baseURL string) *HTTPDoctorClient {
	return &HTTPDoctorClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (c *HTTPDoctorClient) Exists(doctorID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	url := c.baseURL + "/doctors/" + doctorID

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
