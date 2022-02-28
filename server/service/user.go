package service

import (
	"errors"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/utils"
)

type UserService struct{}

func (userService *UserService) Login(u *model.User) (*model.User, error) {
	if !u.CheckEmailIsExist(u.Email) {
		return nil, errors.New("账号不存在")
	}
	res := u.FindByEmail(u.Email)
	if res.Password != u.Password {
		return nil, errors.New("密码错误")
	}
	res.LastLoginTime = time.Now()
	res.Update(res)
	return res, nil
}

func (userService *UserService) Reg(u *model.User) error {
	if u.CheckEmailIsExist(u.Email) {
		return errors.New("邮箱已存在")
	}
	if !utils.Empty(u.Mobile) && u.CheckMobileIsExist(u.Mobile) {
		return errors.New("手机号码已存在")
	}
	return u.Add(u)
}

func (userService *UserService) UpdateUserInfo(u *model.User) error {
	u.UpdateTime = time.Now()
	return u.Update(u)
}

func (userService *UserService) GetUserProfile(uid int) *model.User {
	u := &model.User{}
	return u.FindByUid(uid)
}
