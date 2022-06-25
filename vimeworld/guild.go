package vimeworld

import (
	"context"
	"fmt"
	"net/http"
)

// GuildSearch returns guilds by query.
func (c *Client) GuildSearch(ctx context.Context, query string) ([]*UserGuild, error) {
	var result []*UserGuild
	u := fmt.Sprintf("guild/search?query=%s", query)

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
	u := fmt.Sprintf("guild/get?id=%d", id)

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

// GetGuildByName returns guild by name.
func (c *Client) GetGuildByName(ctx context.Context, name string) (*Guild, error) {
	var result Guild
	u := fmt.Sprintf("guild/get?name=%s", name)

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

// GetGuildByTag returns guild by tag.
func (c *Client) GetGuildByTag(ctx context.Context, tag string) (*Guild, error) {
	var result Guild
	u := fmt.Sprintf("guild/get?tag=%s", tag)

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
