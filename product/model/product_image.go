package model

type ProductImage struct {
	Id int64 `gorm:"primary_key;auto_increment;not null"`
	ImageName string `json:"image_name"`
	ImageSize int64 `json:"image_size"`
	ImageCode string `gorm:"unique_index;not null"`
	ImageUrl string `json:"image_url"`
	ImageProductId int64 `json:"image_product_id"`
}