package vimeworld

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// GuildSearch returns guilds by query.
func (c *Client) GuildSearch(ctx context.Context, query string) ([]*UserGuild, error) {
	var result []*UserGuild

	req, err := c.NewRequest(http.MethodGet, "guild/search", nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("query", query)

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Guild struct.
type Guild struct {
	ID              int                  `json:"id"`
	Name            string               `json:"name"`
	Tag             string               `json:"tag"`
	Color           string               `json:"color"`
	Level           int                  `json:"level"`
	LevelPercentage float64              `json:"levelPercentage"`
	AvatarURL       string               `json:"avatar_url"`
	TotalExp        int                  `json:"totalExp"`
	TotalCoins      int                  `json:"totalCoins"`
	Created         int                  `json:"created"`
	WebInfo         string               `json:"web_info"`
	Perks           map[string]GuildPerk `json:"perks"`
	Members         []GuildMember        `json:"members"`
}

// GuildPerk struct.
type GuildPerk struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// GuildMember struct.
type GuildMember struct {
	User       User   `json:"user"`
	Status     string `json:"status"`
	Joined     int    `json:"joined"`
	GuildCoins int    `json:"guildCoins"`
	GuildExp   int    `json:"guildExp"`
}

// GetGuildByID returns guild by ID.
func (c *Client) GetGuildByID(ctx context.Context, id int) (*Guild, error) {
	var result Guild

	req, err := c.NewRequest(http.MethodGet, "guild/get", nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("id", strconv.Itoa(id))

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGuildByName returns guild by name.
func (c *Client) GetGuildByName(ctx context.Context, name string) (*Guild, error) {
	var result Guild

	req, err := c.NewRequest(http.MethodGet, "guild/get", nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("name", name)

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGuildByTag returns guild by tag.
func (c *Client) GetGuildByTag(ctx context.Context, tag string) (*Guild, error) {
	var result Guild

	req, err := c.NewRequest(http.MethodGet, "guild/get", nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("tag", tag)

	req.URL.RawQuery = urlQuery.Encode()

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
