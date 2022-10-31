package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/marvinhosea/webhook-server/internal/models"
	"net/http"
	"time"
)

type Webhook interface {
	Send(payload *models.Payload) error
}

type Client struct {
	httpClient *http.Client
}

func (c *Client) Send(payload *models.Payload) error {
	body, err := json.Marshal(payload.Data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		"POST",
		payload.WebhookCallbackUrl,
		bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Webhook-Verification", "0014716e-392c-4120-609e-555e295faff5")
	request.Header.Add("User-Agent", "UserAgent-http-client/1.1")
	res, err := c.httpClient.Do(request)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

var _ Webhook = (*Client)(nil)

func NewClient() Webhook {
	client := &http.Client{
		CheckRedirect: nil,
		Timeout:       30 * time.Second,
	}

	return &Client{
		httpClient: client,
	}
}
