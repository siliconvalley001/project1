package handler

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"github.com/siliconvalley001/project1/cart/service"

	"context"
	ex "github.com/siliconvalley001/project1/cartApi/proto"
	"strconv"
)


type CartApi struct{
	C *service.ServiceCart
}

func (c *CartApi)FindAll(context context.Context, resquest *ex.Request, response *ex.Response) error {
	log.Info("接受访问请求")
	if _,ok:=resquest.Get["user_id"];!ok{
		response.StatusCode=500
		return errors.New("参数出错")
	}

	userIdString:=resquest.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	num,err:=strconv.ParseInt(userIdString,10,64)
	if err!=nil{
		return err
	}
	//获取购物车所有商品.
	all,err:=c.C.FindCartAll(
		num,
	)
	//c.C.GetAll(context,&proto.CartFindAll{
	//	UserId: num,
	//},&proto.CartAll{})
	//c.S.GetAll(context,&proto.CartFindAll{
	//	UserId: num,
	//})
	//all,err:=proto.NewCartService().GetAll(context,&proto.CartFindAll{
	//	UserId: num,
	//
	//})
	if err!=nil{}
	return err

	b,err:=json.Marshal(all)
	if err!=nil{
		return err
	}
	response.StatusCode =200
	response.Body = string(b)
	return nil

}





