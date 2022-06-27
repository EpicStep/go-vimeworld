package vimeworld

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// User struct.
type User struct {
	ID              int        `json:"id"`
	Username        string     `json:"username"`
	Level           int        `json:"level"`
	LevelPercentage float64    `json:"levelPercentage"`
	Rank            string     `json:"rank"`
	PlayedSeconds   int64      `json:"playedSeconds"`
	LastSeen        int64      `json:"lastSeen"`
	Guild           *UserGuild `json:"guild"`
}

// UserGuild struct.
type UserGuild struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Tag             string  `json:"tag"`
	Color           string  `json:"color"`
	Level           int     `json:"level"`
	LevelPercentage float64 `json:"levelPercentage"`
	AvatarURL       string  `json:"avatar_url"`
}

// GetUsersByNames return users by names.
func (c *Client) GetUsersByNames(ctx context.Context, names ...string) ([]*User, error) {
	var result []*User
	u := fmt.Sprintf("user/name/%s", strings.Join(names, ","))

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

// GetUsersByIds return users by ids.
func (c *Client) GetUsersByIds(ctx context.Context, ids ...int) ([]*User, error) {
	var result []*User
	var idsStr string

	if len(ids) > 0 {
		idsStr = strconv.Itoa(ids[0])
		for _, id := range ids[1:] {
			idsStr += "," + strconv.Itoa(id)
		}
	}

	u := fmt.Sprintf("user/%s", idsStr)

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

// Friends struct.
type Friends struct {
	User    User   `json:"user"`
	Friends []User `json:"friends"`
}

// GetUserFriends return Friend list of user by ID.
func (c *Client) GetUserFriends(ctx context.Context, id int) (*Friends, error) {
	var result Friends
	u := fmt.Sprintf("user/%d/friends", id)

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

// Session struct.
type Session struct {
	User   User          `json:"user"`
	Online OnlineSession `json:"online"`
}

// OnlineSession struct.
type OnlineSession struct {
	Value   bool   `json:"value"`
	Message string `json:"message"`
	Game    string `json:"game"`
}

// GetUserSession return user session by ID.
func (c *Client) GetUserSession(ctx context.Context, id int) (*Session, error) {
	var result Session
	u := fmt.Sprintf("user/%d/session", id)

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

// UserStats struct.
type UserStats struct {
	User  User                   `json:"user"`
	Stats map[string]interface{} `json:"stats"`
}

// GetUserStats return user stats by ID
func (c *Client) GetUserStats(ctx context.Context, id int, games ...string) (*UserStats, error) {
	var result UserStats
	u := fmt.Sprintf("user/%d/stats", id)

	if len(games) > 0 {
		u += "?games=" + strings.Join(games, ",")
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

// UserAchievements struct.
type UserAchievements struct {
	User         User              `json:"user"`
	Achievements []UserAchievement `json:"achievements"`
}

// UserAchievement struct.
type UserAchievement struct {
	ID   int   `json:"ID"`
	Time int64 `json:"time"`
}

// GetUserAchievements return user achievements by ID.
func (c *Client) GetUserAchievements(ctx context.Context, id int) (*UserAchievements, error) {
	var result UserAchievements
	u := fmt.Sprintf("user/%d/achievements", id)

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

// UserLeaderboards struct.
type UserLeaderboards struct {
	User         User              `json:"user"`
	Leaderboards []UserLeaderboard `json:"leaderboards"`
}

// UserLeaderboard struct.
type UserLeaderboard struct {
	Type  string `json:"type"`
	Sort  string `json:"sort"`
	Place int    `json:"place"`
}

// GetUserLeaderboards returns user leaderboard by ID.
func (c *Client) GetUserLeaderboards(ctx context.Context, id int) (*UserLeaderboards, error) {
	var result UserLeaderboards
	u := fmt.Sprintf("user/%d/leaderboards", id)

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

// UserMatches struct.
type UserMatches struct {
	User    User `json:"user"`
	Request struct {
		Count  int `json:"count"`
		Offset int `json:"offset"`
		Size   int `json:"size"`
	} `json:"request"`
	Matches []UserMatch `json:"matches"`
}

// UserMatch struct.
type UserMatch struct {
	ID       string   `json:"id"`
	Game     string   `json:"game"`
	Map      MatchMap `json:"map"`
	Date     int      `json:"date"`
	Duration int      `json:"duration"`
	Players  int      `json:"players"`
	Win      bool     `json:"win"`
	State    int      `json:"state"`
}

// GetUserMatchesAfter return matches after id.
func (c *Client) GetUserMatchesAfter(ctx context.Context, id int, after string, count int) (*UserMatches, error) {
	var result UserMatches
	u := fmt.Sprintf("user/%d/matches?after=%s", id, after)

	if count > 0 {
		u += "&count=" + strconv.Itoa(count)
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

// GetUserMatchesBefore return matches before id.
func (c *Client) GetUserMatchesBefore(ctx context.Context, id int, before string, count int) (*UserMatches, error) {
	var result UserMatches
	u := fmt.Sprintf("user/%d/matches?before=%s", id, before)

	if count > 0 {
		u += "&count=" + strconv.Itoa(count)
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

// GetUserMatchesOffset return matches by offset.
func (c *Client) GetUserMatchesOffset(ctx context.Context, id, offset, count int) (*UserMatches, error) {
	var result UserMatches
	u := fmt.Sprintf("user/%d/matches?offset=%d", id, offset)

	if count > 0 {
		u += "&count=" + strconv.Itoa(count)
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

// UserSession struct.
type UserSession struct {
	User
	Online OnlineSession `json:"online"`
}

// GetUsersSessions return user sessions by ids.
func (c *Client) GetUsersSessions(ctx context.Context, ids ...int) ([]*UserSession, error) {
	var err error
	var result []*UserSession
	var req *http.Request

	if len(ids) > 50 {
		u := "user/session"

		req, err = c.NewRequest(http.MethodPost, u, ids)
		if err != nil {
			return nil, err
		}
	} else {
		var idsStr string

		if len(ids) > 0 {
			idsStr = strconv.Itoa(ids[0])
			for _, id := range ids[1:] {
				idsStr += "," + strconv.Itoa(id)
			}
		}

		u := fmt.Sprintf("user/session/%s", idsStr)

		req, err = c.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}
	}

	_, err = c.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
