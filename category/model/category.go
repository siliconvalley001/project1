package model

import "time"

type Category struct {
	Id     int64 `gorm:"primary_key;auto_increment " json:"id"`
	Parent int64 `json:"parent" `
	Name   string  `gorm:"unique_index;not null"`
	Level  uint32 `json:"level"`
	Image  string `json:"image"`
	Desc   string `json:"desc"`
	sort   int64 `json:"sort"`
	Create time.Time
}
