package service

import "github.com/siliconvalley001/project1/cart/model"

func (c *ServiceCart)AddCart(cart *model.Cart)(int64,error){
	return c.cart.CreateCart(cart)

}

func (c *ServiceCart)DelCart(cartid int64)error{
	return c.cart.DelCartById(cartid)
}


func (c *ServiceCart)UpdateCart(cart *model.Cart)error{
	return c.cart.UpdateCart(cart)
}

func (c *ServiceCart)FindCartByID(cartid int64)(*model.Cart,error){
	return c.cart.FindCartById(cartid)
}

func (c *ServiceCart)FindCartAll(userid int64)([]model.Cart,error)  {
	return c.cart.FindAll(userid)
}


func (c *ServiceCart)CleanCart(userid int64)error{
	return c.cart.CleanCart(userid)
}


func (c *ServiceCart)DecrNum(cartid int64,num int64)error{
	return  c.cart.DecrNum(cartid,num)
}


func (c *ServiceCart)IncrNum(cartid int64,num int64)error{
	return c.cart.IncrNum(cartid,num)
}