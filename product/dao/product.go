package dao

import "github.com/siliconvalley001/project1/product/model"

func (d *Dao)FindProductById(id int64)(product *model.Product,err error){
	product=&model.Product{}
	return product,d.engine.Preload("Images").Preload("Seo").Preload("Size").First(product,id).Error

}

func (d *Dao)CreateProduct(product *model.Product)(int64,error){
	return product.Id,d.engine.Create(product).Error
}


func (d *Dao)DeleteProductById(productid int64)error{
	//开启事务
	tx:=d.engine.Begin()
	defer func() {
		if r:=recover();r!=nil{
			tx.Rollback()
		}
	}()
	if tx.Error!=nil{
		return tx.Error
	}
	//delete
	if err:=tx.Unscoped().Where(`id=?`,productid).Delete(&model.Product{}).Error;err!=nil{
		tx.Rollback()
		return err
	}
	if err:=tx.Unscoped().Where(`size_product_id=?`,productid).Delete(&model.ProductSize{}).Error;err!=nil{
		tx.Rollback()
		return err
	}
	if err:=tx.Unscoped().Where(`seo_product_id=?`,productid).Delete(&model.ProductSeo{}).Error;err!=nil{
		tx.Rollback()
		return err
	}
	if err:=tx.Unscoped().Where(`image_product_id=?`,productid).Delete(&model.ProductImage{}).Error;err!=nil{
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}


func(d *Dao)UpdateProduct(product *model.Product)error{

	return d.engine.Model(product).Update(product).Error

}


func (d *Dao)FindAll()(productall []model.Product,err error ){
	return productall,d.engine.Preload("Images").Preload("Seo").Preload("Size").Find(&productall).Error

}


