package vimeworld

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// Match struct.
type Match struct {
	ID       string   `json:"id"`
	Game     string   `json:"game"`
	Map      MatchMap `json:"map"`
	Date     int64    `json:"date"`
	Duration int      `json:"duration"`
	Players  int      `json:"players"`
}

// MatchMap struct.
type MatchMap struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Teams         int    `json:"teams"`
	PlayersInTeam int    `json:"playersInTeam"`
}

// GetMatchLatest returns latest matches.
func (c *Client) GetMatchLatest(ctx context.Context) ([]*Match, error) {
	var result []*Match

	req, err := c.NewRequest(http.MethodGet, "match/latest", nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMatchListBefore returns matches before ID.
func (c *Client) GetMatchListBefore(ctx context.Context, before string, count int) ([]*Match, error) {
	var result []*Match

	req, err := c.NewRequest(http.MethodGet, "match/list", nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("before", before)
	if count > 0 {
		urlQuery.Set("count", strconv.Itoa(count))
	}

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMatchListAfter returns matches after ID.
func (c *Client) GetMatchListAfter(ctx context.Context, after string, count int) ([]*Match, error) {
	var result []*Match

	req, err := c.NewRequest(http.MethodGet, "match/list", nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("after", after)
	if count > 0 {
		urlQuery.Set("count", strconv.Itoa(count))
	}

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
