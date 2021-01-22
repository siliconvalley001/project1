package dao

import "github.com/siliconvalley001/project1/category/model"

func (d *Dao)InitTable()error{
	return d.engine.CreateTable(&model.Category{}).Error
}
func (d *Dao)CreateCategory(category *model.Category)(id int64,err error){
	return category.Id, d.engine.Create(category).Error
}
func (d *Dao)UpdateCategory(category *model.Category)error{
	return d.engine.Model(category).Update(&category).Error

}
func (d *Dao)DeleteCategoryById(id int64)error{
	return d.engine.Where(`id=?`, id).Delete(&model.Category{}).Error
}
func (d *Dao)FindCategoryByName(name string) (user *model.Category, err error) {
	category:= &model.Category{}
	return user, d.engine.Where(`name= ?`, name).Find(category).Error
}
func (d *Dao)FindCategoryByID(id int64)(cate *model.Category,err error){
	cate=&model.Category{}
	return cate, d.engine.First(cate, id).Find(cate).Error
}
func (d *Dao)FindCategoryByLevel(level uint32)(allcate []model.Category,err error){
	return allcate,d.engine.Find(`level= ?`,level).Find(&allcate).Error

}
func (d *Dao)FindCategoryByParent(parent int64)(allcate []model.Category,err error){
	return allcate,d.engine.Find(`parent= ?`,parent).Find(&allcate).Error

}

func (d *Dao)FindAllCategory()(all []model.Category,err error){
	return all, d.engine.Find(&all).Error
}