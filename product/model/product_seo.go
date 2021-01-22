package model


type ProductSeo struct {
	Id int64 `gorm:"primary_key;auto_increment;not null"`
	SeoTitle string `json:"seo_title"`
	SeoKeywords string `json:"seo_keywords"`
	SeoDesc string `json:"seo_desc"`
	SeoCode string `json:"seo_code"`
	SeoProductId int64 `json:"seo_product_id"`
}
