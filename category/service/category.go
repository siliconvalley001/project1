package service

import "github.com/siliconvalley001/project1/category/model"

func (s *Service_Category)AddCategory(cate *model.Category)(int64,error){
	return s.dao.CreateCategory(cate)

}


func(s *Service_Category)DelCategory(id int64)error{
	return s.dao.DeleteCategoryById(id)
}

func (s *Service_Category)UpdateCategoryByID(id int64)error{
	return s.dao.DeleteCategoryById(id)
}
func (s *Service_Category)FindCategoryByName(catename string)(cate *model.Category,err error){
	return s.dao.FindCategoryByName(catename)
}
func (s *Service_Category)FindCategoryById(id int64)(cate *model.Category,err error){
	return s.dao.FindCategoryByID(id)
}
func (s *Service_Category)FindCategoryByLevel(level uint32)(all []model.Category,err error){
	return s.dao.FindCategoryByLevel(level)
}
func (s *Service_Category)FindCategoryByParent(parent int64)(all []model.Category,err error){
	return s.dao.FindCategoryByParent(parent)
}
func (s *Service_Category)FindCategoryAll()(all []model.Category,err error){
	return s.dao.FindAllCategory()
}
