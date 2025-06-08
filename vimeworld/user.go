package vimeworld

import (
	"context"
	"net/http"
	"net/url"
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
	u := "user/name/" + url.PathEscape(strings.Join(names, ","))

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

// GetUsersByIDs return users by ids.
func (c *Client) GetUsersByIDs(ctx context.Context, ids ...int) ([]*User, error) {
	var result []*User
	var idsStr string

	if len(ids) > 0 {
		idsStr = strconv.Itoa(ids[0])
		for _, id := range ids[1:] {
			idsStr += "," + strconv.Itoa(id)
		}
	}

	u := "user/" + url.PathEscape(idsStr)

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
	u := "user/" + strconv.Itoa(id) + "/friends"

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
	u := "user/" + strconv.Itoa(id) + "/session"

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
	User  User           `json:"user"`
	Stats map[string]any `json:"stats"`
}

// GetUserStats return user stats by ID
func (c *Client) GetUserStats(ctx context.Context, id int, games ...string) (*UserStats, error) {
	var result UserStats
	u := "user/" + strconv.Itoa(id) + "/stats"

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	if len(games) > 0 {
		urlQuery.Set("games", strings.Join(games, ","))
	}

	req.URL.RawQuery = urlQuery.Encode()

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
	u := "user/" + strconv.Itoa(id) + "/achievements"

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
	u := "user/" + strconv.Itoa(id) + "/leaderboards"

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
	User    User               `json:"user"`
	Request UserMatchesRequest `json:"request"`
	Matches []UserMatch        `json:"matches"`
}

// UserMatchesRequest ...
type UserMatchesRequest struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Size   int `json:"size"`
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
	u := "user/" + strconv.Itoa(id) + "/matches"

	req, err := c.NewRequest(http.MethodGet, u, nil)
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

	return &result, nil
}

// GetUserMatchesBefore return matches before id.
func (c *Client) GetUserMatchesBefore(ctx context.Context, id int, before string, count int) (*UserMatches, error) {
	var result UserMatches
	u := "user/" + strconv.Itoa(id) + "/matches"

	req, err := c.NewRequest(http.MethodGet, u, nil)
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

	return &result, nil
}

// GetUserMatchesOffset return matches by offset.
func (c *Client) GetUserMatchesOffset(ctx context.Context, id, offset, count int) (*UserMatches, error) {
	var result UserMatches
	u := "user/" + strconv.Itoa(id) + "/matches"

	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	urlQuery := url.Values{}
	urlQuery.Set("offset", strconv.Itoa(offset))
	if count > 0 {
		urlQuery.Set("count", strconv.Itoa(count))
	}

	req.URL.RawQuery = urlQuery.Encode()

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
		req, err = c.NewRequest(http.MethodPost, "user/session", ids)
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

		req, err = c.NewRequest(http.MethodGet, "user/session/"+idsStr, nil)
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
