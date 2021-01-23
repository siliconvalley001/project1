package handler

import (
	"github.com/siliconvalley001/project1/category/common"
	"github.com/siliconvalley001/project1/category/model"
	ex "github.com/siliconvalley001/project1/category/proto"
	"github.com/siliconvalley001/project1/category/service"
	"context"
	"log"
)

type Category struct {
	srv service.Service_Category
}

func (c *Category) CreateCategory(context context.Context, req *ex.CategoryRequest, resp *ex.CreateCategoryResponse) error {
	cate := &model.Category{}
	err := common.Swap(req, cate)
	if err != nil {
		return err
	}
	id, err := c.srv.AddCategory(cate)
	if err != nil {
		return err
	}
	resp.CategoryId = id
	resp.Msg = "分类添加成功"
	return nil

}

func (c *Category) UpdateCategory(con context.Context, req *ex.CategoryRequest,resp  *ex.UpdateCategoryResponse) error {
	cate := &model.Category{}
	err := common.Swap(req, cate)
	if err != nil {
		return err
	}
	err= c.srv.UpdateCategoryByID(cate.Id)

	if err!=nil{
		return err
	}

	resp.Msg="分类更新成功"
	return nil
}
 func (c *Category)DeleteCategory(con context.Context,req *ex.DeleteCategoryResquest,resp *ex.DeleteCategoryResponse) error {
 	cate:=&model.Category{}
	 err := common.Swap(req, cate)
	 if err != nil {
		 return err
	 }
	err=c.srv.DelCategory(req.CategoryId)
	if err!=nil{
		return err
	}
	resp.Msg="分类删除成功"
	return nil
 }
func (c *Category)FindCategoryByName(con context.Context, resq *ex.CategoryByNameResquest, resp *ex.CategoryResponse) error{

	cate,err:=c.srv.FindCategoryByName(resq.CategoryName)
	if err!=nil{
		return err
	}
	return common.Swap(cate,resp)

}
func (c *Category)FindCategoryById(ctx context.Context, resquest *ex.CategoryByIdResquest, response *ex.CategoryResponse) error{
	cate, err := c.srv.FindCategoryById(resquest.CategoryId)
	if err!=nil{
		return err
	}
	return common.Swap(cate,response)
}

func (c *Category)FindCategoryByLevel(ctx context.Context, resquest *ex.CategoryByLevelResquest,response *ex.FindAllResponse) error{
	cate,err:=c.srv.FindCategoryByLevel(resquest.CategoryLevel)
	if err!=nil{
		return err
	}
	cateToResponse(cate,response)
	return nil
}

func cateToResponse(cate []model.Category,response *ex.FindAllResponse){
	for _,value:=range cate{
		c:=&ex.CategoryResponse{}
		err:=common.Swap(value,c)
		if err!=nil{
			log.Fatal(err)
			break
		}
		response.All=append(response.All,c)
	}

}

func  (c *Category)FindCategoryByParent(ctx context.Context, resquest *ex.CategoryByParentResquest, response *ex.FindAllResponse) error {
	all,err:=c.srv.FindCategoryByParent(resquest.ParentId)
	if err!=nil{
		return err
	}
	cateToResponse(all,response)
	return nil
	
}

func (c *Category)FindAllCategory(ctx context.Context,  resquest*ex.FindAllResquest, response *ex.FindAllResponse) error{
	all,err:=c.srv.FindCategoryAll()
	if err!=nil{
		return err
	}
	cateToResponse(all,response)
	return nil
}
