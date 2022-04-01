package models

import (
	"time"
)

type Host struct {
	ID         int64     `json:"id" db:"post_id"`
	Region     string    `json:"region" db:"region"`
	Hostname   string    `json:"hostname" db:"hostname"`
	Publicip   string    `json:"publicip" db:"publicip"`
	Privateip  string    `json:"privateip" db:"privateip"`
	Os         string    `json:"os" db:"os"`
	Status     string    `json:"status" db:"status"`
	Remark     string    `json:"remark" db:"remark"`
	DeleteTime time.Time `json:"delete_time" db:"delete_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
