package dao

import (
	"github.com/siliconvalley001/project1/user/model"
)

func (d *Dao) InitUserTable() error {
	return d.engine.CreateTable(&model.User{}).Error
}

func (d *Dao) FindUserByNickName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, d.engine.Where(`nick_name= ?`, name).Find(user).Error
}

func (d *Dao) FindUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	return user, d.engine.First(user, id).Find(user).Error

}

func (d *Dao) CreateUser(user *model.User) (id int64, err error) {


	return user.Id, d.engine.Create(user).Error
}

func (d *Dao) DelUserByID(id int64) error {
	return d.engine.Where(`id=?`, id).Delete(&model.User{}).Error

}

func (d *Dao) UpdateUser(user *model.User) error {
	return d.engine.Model(user).Update(&user).Error
}

func (d *Dao) FindAll() (userall []model.User, err error) {
	return userall, d.engine.Find(&userall).Error
}

