package vimeworld

import (
	"context"
	"net/http"
	"net/url"
)

// Game struct.
type Game struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	GlobalStats []string          `json:"global_stats"`
	SeasonStats GameSessionsStats `json:"season_stats"`
}

// GameSessionsStats ...
type GameSessionsStats struct {
	Monthly []string `json:"monthly"`
	Manual  []string `json:"manual"`
}

// GetGames returns games.
func (c *Client) GetGames(ctx context.Context) ([]*Game, error) {
	var result []*Game

	req, err := c.NewRequest(http.MethodGet, "misc/games", nil)
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

	req, err := c.NewRequest(http.MethodGet, "misc/maps", nil)
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

	req, err := c.NewRequest(http.MethodGet, "misc/achievements", nil)
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
	u := "misc/token/" + url.PathEscape(token)

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
