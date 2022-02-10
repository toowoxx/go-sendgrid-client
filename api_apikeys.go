package sendgrid

import "github.com/sendgrid/rest"

type APIScope = string
type APIKey struct {
	Name   string     `json:"name"`
	Scopes []APIScope `json:"scopes"`
}

func (c *Client) CreateAPIKey(apiKeyDetails APIKey) error {
	return c.Request("/v3/api_keys", rest.Post, apiKeyDetails, nil)
}
