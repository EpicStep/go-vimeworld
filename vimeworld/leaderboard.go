package vimeworld

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// AvailableLeaderboard struct.
type AvailableLeaderboard struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	MaxSize     int      `json:"max_size"`
	Sort        []string `json:"sort"`
	Season      string   `json:"season"`
}

// GetLeaderboardList returns available leaderboards.
func (c *Client) GetLeaderboardList(ctx context.Context) ([]*AvailableLeaderboard, error) {
	var result []*AvailableLeaderboard

	req, err := c.NewRequest(http.MethodGet, "leaderboard/list", nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Leaderboard struct.
type Leaderboard struct {
	Meta    LeaderboardResponseMeta `json:"leaderboard"`
	Records []any                   `json:"records"`
}

// LeaderboardResponseMeta struct.
type LeaderboardResponseMeta struct {
	Type    string `json:"type"`
	Sort    string `json:"sort"`
	Offset  int    `json:"offset"`
	Size    int    `json:"size"`
	MaxSize int    `json:"max_size"`
}

// GetLeaderboard returns leaderboard.
func (c *Client) GetLeaderboard(ctx context.Context, boardType, sort string, size, offset int) (*Leaderboard, error) {
	var result Leaderboard
	u := "leaderboard/get/" + url.PathEscape(boardType)

	if sort != "" {
		u += "/" + url.PathEscape(sort)
	}

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	if size > 0 {
		urlQuery.Set("size", strconv.Itoa(size))
	}

	if offset > 0 {
		urlQuery.Set("offset", strconv.Itoa(offset))
	}

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
