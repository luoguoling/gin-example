package models

type Community struct {
	CommunityID   int64  `json:"community_id"`
	CommunityName string `json:"community_name"`
	Introduction  string `json:"introduction"`
}
