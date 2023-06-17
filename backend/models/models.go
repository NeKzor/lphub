package models

import (
	"time"
)

type User struct {
	SteamID     string    `json:"steam_id"`
	UserName    string    `json:"user_name"`
	AvatarLink  string    `json:"avatar_link"`
	CountryCode string    `json:"country_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserShort struct {
	SteamID  string `json:"steam_id"`
	UserName string `json:"user_name"`
}

type Map struct {
	ID          int    `json:"id"`
	GameName    string `json:"game_name"`
	ChapterName string `json:"chapter_name"`
	MapName     string `json:"map_name"`
	Data        any    `json:"data"`
}

type MapShort struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MapSummary struct {
	Description    string            `json:"description"`
	Showcase       string            `json:"showcase"`
	CategoryScores MapCategoryScores `json:"category_scores"`
	Rating         float32           `json:"rating"`
	Routers        []string          `json:"routers"`
	History        []MapHistory      `json:"history"`
}

type MapCategoryScores struct {
	CM          int `json:"cm"`
	NoSLA       int `json:"no_sla"`
	InboundsSLA int `json:"inbounds_sla"`
	Any         int `json:"any"`
}

type MapRecords struct {
	Records any `json:"records"`
}

type MapHistory struct {
	RunnerName string    `json:"runner_name"`
	ScoreCount int       `json:"score_count"`
	Date       time.Time `json:"date"`
}

type UserRanking struct {
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	TotalScore int    `json:"total_score"`
}

type Game struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Chapter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RecordSP struct {
	RecordID   int       `json:"record_id"`
	Placement  int       `json:"placement"`
	UserID     string    `json:"user_id"`
	UserName   string    `json:"user_name"`
	UserAvatar string    `json:"user_avatar"`
	ScoreCount int       `json:"score_count"`
	ScoreTime  int       `json:"score_time"`
	DemoID     string    `json:"demo_id"`
	RecordDate time.Time `json:"record_date"`
}

type RecordMP struct {
	RecordID      int       `json:"record_id"`
	Placement     int       `json:"placement"`
	HostID        string    `json:"host_id"`
	HostName      string    `json:"host_name"`
	HostAvatar    string    `json:"host_avatar"`
	PartnerID     string    `json:"partner_id"`
	PartnerName   string    `json:"partner_name"`
	PartnerAvatar string    `json:"partner_avatar"`
	ScoreCount    int       `json:"score_count"`
	ScoreTime     int       `json:"score_time"`
	HostDemoID    string    `json:"host_demo_id"`
	PartnerDemoID string    `json:"partner_demo_id"`
	RecordDate    time.Time `json:"record_date"`
}

type PlayerSummaries struct {
	SteamId                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaName              string `json:"personaname"`
	LastLogOff               int    `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`

	CommentPermission int    `json:"commentpermission"`
	RealName          string `json:"realname"`
	PrimaryClanId     string `json:"primaryclanid"`
	TimeCreated       int    `json:"timecreated"`
	LocCountryCode    string `json:"loccountrycode"`
	LocStateCode      string `json:"locstatecode"`
	LocCityId         int    `json:"loccityid"`
	GameId            int    `json:"gameid"`
	GameExtraInfo     string `json:"gameextrainfo"`
	GameServerIp      string `json:"gameserverip"`
}
