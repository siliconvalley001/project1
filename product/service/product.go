package service

import "github.com/siliconvalley001/project1/product/model"

func (s *ServiceProduct)AddProduct(product *model.Product)(int64,error){
	return s.product.CreateProduct(product)

}

func (s *ServiceProduct)DelProduct(productid int64)error{
	return s.product.DeleteProductById(productid)
}

func (s *ServiceProduct)UpdateProduct(product *model.Product)error{
	return s.product.UpdateProduct(product)
}

func (s *ServiceProduct)FindProductById(productid int64)(*model.Product,error){
	return s.product.FindProductById(productid)
}

func (s *ServiceProduct)FindAllProduct()([]model.Product,error){
	return s.product.FindAll()
}
