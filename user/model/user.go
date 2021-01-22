package model

import "time"

type User struct {
	Id int64 `gorm:"primary_key;not_null,auto_increment"`
	NickName string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Name string  `gorm:"not null"`
	OpenId string `gorm:"not null"`
	Integral  int64 `gorm:"not null"`
	Avatar  string  `gorm:"not null"`
	Tel string `gorm:"not null"`
	Email string `gorm:"not null"`
	CreateTime time.Time
	Sex int64 `gorm:"not null"`
	Del int32 `gorm:"not null"`
}
