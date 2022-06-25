package vimeworld

import (
	"context"
	"fmt"
	"net/http"
)

// Game struct.
type Game struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	GlobalStats []string `json:"global_stats"`
	SeasonStats struct {
		Monthly []string `json:"monthly"`
		Manual  []string `json:"manual"`
	} `json:"season_stats"`
}

// GetGames returns games.
func (c *Client) GetGames(ctx context.Context) ([]*Game, error) {
	var result []*Game
	u := "misc/games"

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

// Maps struct.
type Maps map[string]map[string]Map

// Map struct.
type Map struct {
	Name          string `json:"name"`
	Teams         int    `json:"teams"`
	PlayersInTeam int    `json:"playersInTeam"`
}

// GetMaps returns maps.
func (c *Client) GetMaps(ctx context.Context) (*Maps, error) {
	var result Maps
	u := "misc/maps"

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

// Achievements struct.
type Achievements map[string][]Achievement

// Achievement struct.
type Achievement struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Reward      int      `json:"reward"`
	Description []string `json:"description"`
}

// GetAchievements returns achievements.
func (c *Client) GetAchievements(ctx context.Context) (*Achievements, error) {
	var result Achievements
	u := "misc/achievements"

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

// Token struct.
type Token struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
	Type  string `json:"type"`
	Limit int    `json:"limit"`
	Owner *User  `json:"owner"`
}

// GetTokenInfo returns token info by token.
func (c *Client) GetTokenInfo(ctx context.Context, token string) (*Token, error) {
	var tokenResponse Token
	u := fmt.Sprintf("misc/token/%s", token)

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, &tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
