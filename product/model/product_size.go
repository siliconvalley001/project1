package model


type ProductSize struct {
	Id int64 `gorm:"primary_key;auto_increment;not null"`
	SizeName string `json:"size_name"`
	SizeCode string `json:"size_code"`
	SizeProductId int64 `json:"size_product_id"`
	

}
