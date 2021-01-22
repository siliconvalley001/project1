package dao

import (
	"github.com/siliconvalley001/project1/cart/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (d *Dao) FindCartById(cart_id int64) (cart *model.Cart, err error) {

	cart = &model.Cart{}
	return cart, d.engine.First(cart, cart_id).Error
}

func (d *Dao) CreateCart(cart *model.Cart) (int64, error) {
	db := d.engine.FirstOrCreate(cart, model.Cart{ProductId: cart.ProductId, SizeId: cart.SizeId})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected <= 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.Id, nil
}
func (d *Dao) DelCartById(cart_id int64) error {
	return d.engine.Where(`cart_id=?`, cart_id).Delete(&model.Cart{}).Error
}

func (d *Dao) UpdateCart(cart *model.Cart) error {
	return d.engine.Model(cart).Update(cart).Error
}

func (d *Dao) FindAll(user_id int64) (all []model.Cart, err error) {
	return all, d.engine.Where(`user_id=?`, user_id).Find(&all).Error
}

func (d *Dao) CleanCart(user_id int64) error {
	return d.engine.Where(`user_id=?`, user_id).Delete(&model.Cart{}).Error
}

func (d *Dao) IncrNum(cartid int64, num int64) error {
	cart := &model.Cart{
		Id: cartid,
	}
	return d.engine.Model(cart).UpdateColumn("num", gorm.Expr(`num=?`, num)).Error
}

func (d *Dao) DecrNum(cartid int64, num int64) error {
	cart := &model.Cart{
		Id: cartid,
	}
	db:=d.engine.Model(cart).Where(`num> ?`,num).UpdateColumn("num", gorm.Expr(`num=?`, num))

	if db.Error!=nil{
		return db.Error
	}
	return nil


}
