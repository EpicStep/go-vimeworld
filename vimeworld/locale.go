package vimeworld

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// Locale struct.
type Locale struct {
	Games     map[string]GameLocale     `json:"games"`
	GameStats map[string]map[string]any `json:"game_stats"`
	Ranks     map[string]RankLocale     `json:"ranks"`
}

// GameLocale struct.
type GameLocale struct {
	Name string `json:"name"`
}

// RankLocale struct.
type RankLocale struct {
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}

// GetLocaleByName returns locale by name.
func (c *Client) GetLocaleByName(ctx context.Context, name string, parts ...string) (*Locale, error) {
	var result Locale
	u := "locale/" + url.PathEscape(name)

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	if len(parts) > 0 {
		urlQuery.Set("parts", strings.Join(parts, ","))
	}

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
