package handler

import (
	"github.com/siliconvalley001/project1/cart/common"
	"github.com/siliconvalley001/project1/cart/model"
	"github.com/siliconvalley001/project1/cart/service"
	"context"
	ex "github.com/siliconvalley001/project1/cart/proto"

)

type Cart struct{
	srv *service.ServiceCart
}


func (c *Cart)	AddCart(ctx context.Context, resquest *ex.CartInfo, response *ex.ResponseAdd) error{
	cart:=&model.Cart{}
	//common.Swap(resquest,cart)
	common.Swap(resquest,cart)

	if _,err:=c.srv.AddCart(cart);err!=nil{
		return err
	}
	return nil

}

func (c *Cart)CleanCart(ctx context.Context,resquest *ex.Clean, response *ex.Response) error{
	if err:=c.srv.CleanCart(resquest.UserId);err!=nil{
		return err
	}
	response.Meg="购物车清空成功"
	return nil

}
func (c *Cart)Incr(ctx context.Context,resquest *ex.Item, response *ex.Response) error{
	if err:=c.srv.IncrNum(resquest.Id,resquest.ChangeNum);err!=nil{
		return err
	}
	response.Meg="购物车新增成功"
	return nil
}
func (c *Cart)Decr(ctx context.Context, resquest *ex.Item,response *ex.Response) error{
	if err:=c.srv.DecrNum(resquest.Id,resquest.ChangeNum);err!=nil{
		return err
	}
	response.Meg="购物车减少成功"
	return nil

}
func (c *Cart)DeleteItemByID(ctx context.Context,resquest *ex.CartID,response *ex.Response) error{
	if err:=c.srv.DelCart(resquest.Id);err!=nil{
		return err
	}
	response.Meg="购物车删除成功"
	return nil
}
func (c *Cart)GetAll(ctx context.Context, resquest *ex.CartFindAll,response *ex.CartAll) error{
	all,err:=c.srv.FindCartAll(resquest.UserId)
	if err!=nil{
		return err
	}
	for _,v:=range all{
		cart:=&ex.CartInfo{}
		if err:=common.Swap(v,cart);err!=nil{
			return err
		}
		response.CartInfo=append(response.CartInfo,cart)



	}
	return nil

}



