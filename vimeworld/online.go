package vimeworld

import (
	"context"
	"net/http"
)

// Online struct.
type Online struct {
	Total     int            `json:"total"`
	Separated map[string]int `json:"separated"`
}

// GetOnline return online of a server.
func (c *Client) GetOnline(ctx context.Context) (*Online, error) {
	var result Online
	u := "online"

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// OnlineStream struct.
type OnlineStream struct {
	Title    string `json:"title"`
	Owner    string `json:"owner"`
	Viewers  int    `json:"viewers"`
	URL      string `json:"url"`
	Duration int    `json:"duration"`
	Platform string `json:"platform"`
	User     User   `json:"user"`
}

// GetOnlineStreams returns online streams.
func (c *Client) GetOnlineStreams(ctx context.Context) ([]*OnlineStream, error) {
	var result []*OnlineStream
	u := "online/streams"

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetOnlineStaff returns online staff.
func (c *Client) GetOnlineStaff(ctx context.Context) ([]*User, error) {
	var result []*User
	u := "online/staff"

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
