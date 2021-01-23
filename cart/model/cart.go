package model



type Cart struct {
	Id int64 `gorm:"primary_key;auto_increment;not null"`
	ProductId int64 `gorm:"not null" json:"product_id"`
	Num  int64 `gorm:"not null" json:"num"`
	SizeId int64 `gorm:"not null" json:"size_id"`
	UserId int64 `gorm:"not null" json:"user_id"`
}



