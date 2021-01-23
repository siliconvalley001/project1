package model


type Product struct {
	Id int64 `gorm:"primary_key;auto_increment;not null"`
	Name string
	Digitnum string `gorm:"unique_index;not null"`
	Price int64
	Desc string
	Images []ProductImage `gorm:"ForeignKey:ImagesProductId"`
	Size []ProductSize `gorm:"ForeignKey:SizeProductId"`
	Seo  ProductSeo  `gorm:"ForeignKey:SeoProductId"`
}


