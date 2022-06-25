package vimeworld

import (
	"context"
	"fmt"
	"net/http"
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
	u := "leaderboard/list"

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

// Leaderboard struct.
type Leaderboard struct {
	Meta    LeaderboardResponseMeta `json:"leaderboard"`
	Records []interface{}           `json:"records"`
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
	u := fmt.Sprintf("leaderboard/get/%s", boardType)

	if sort != "" {
		u += "/" + sort
	}

	if size > 0 {
		u += "?size=" + strconv.Itoa(size)
	}

	if offset > 0 {
		if size > 0 {
			u += "&offset=" + strconv.Itoa(offset)
		} else {
			u += "?offset=" + strconv.Itoa(offset)
		}
	}

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
