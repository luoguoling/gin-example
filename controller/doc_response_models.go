package controller

import "time"

type _Community struct {
	ID   uint64 `json:"id"`   //帖子分类id
	Name string `json:"name"` //帖子分类名称
}

// _CommunityDetail 帖子分类响应数据
type _CommunityDetail struct {
	ID           uint64    `json:"id"`           //id值
	Name         string    `json:"name"`         //帖子分类名称
	Introduction string    `json:"introduction"` //分类介绍
	CreateTime   time.Time `json:"create_time"`  //创建时间
}
