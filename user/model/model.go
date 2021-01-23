package model

import (
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"github.com/siliconvalley001/project1/user/setting"
	"fmt"
)
var (
	DB *gorm.DB
	err error
)

type AllConfig struct {

}

func init(){
	if _,err:=setting.InitSetting();err!=nil{
		panic(err)
	}
	fmt.Println(setting.Con.Mys.Name)
	DB,err=gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		setting.Con.Mys.Name,
		setting.Con.Mys.PassWord,
		setting.Con.Mys.Host,
		setting.Con.Mys.Port,
		setting.Con.Mys.DBName,
	))
	if err!=nil{
		panic(err)

	}
	DB.SingularTable(true)

	DB.DB().SetMaxOpenConns(setting.Con.Mys.MaxOpenConns)

	DB.DB().SetMaxIdleConns(setting.Con.Mys.MaxIdleConns)

	DB.AutoMigrate(&User{})


}
