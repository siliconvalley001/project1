package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/siliconvalley001/project1/user/model"
)

func (s *Service_User) AddUser(user *model.User) (id int64, err error) {
	user.Password=generatePwd(user.Password)
	return s.dao.CreateUser(user)

}

func (s *Service_User) DelUser(userid int64) (err error) {
	return s.dao.DelUserByID(userid)

}

func (s *Service_User) UpdateUser(user *model.User, IsChangePWD bool) (err error) {
	if IsChangePWD{
		user.Password=generatePwd(user.Password)

	}
	return s.dao.UpdateUser(user)
}

func (s *Service_User) FindUserByNickName(nickname string) (user *model.User, err error) {

	return s.dao.FindUserByNickName(nickname)
}

func (s *Service_User) CheckPassWord(nickname string,password string)(err error) {
	Hash:=generatePwd(password)
	userModel,err:=s.FindUserByNickName(nickname)
	if err!=nil{
		return err
	}
	if userModel.Password!=Hash{
		return errors.New("用户名或密码错误")
	}
	return


}


func generatePwd(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	PWD := hash.Sum(nil)
	return hex.EncodeToString(PWD)
}

func ValiatePwd(password string, hashpassword string) bool {
	return generatePwd(password) == hashpassword
}