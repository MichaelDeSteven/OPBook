package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

// User
type User struct {
	Id            int       `gorm:"primaryKey;auto;" json:"id"`
	Nickname      string    `gorm:"" json:"nickname"`                            // 昵称
	Password      string    `gorm:"" json:"password"`                            // 密码
	Email         string    `gorm:"" json:"email"`                               // 邮箱，唯一
	Mobile        string    `gorm:"" json:"mobile"`                              // 手机号码，唯一
	Avatar        string    `gorm:"" json:"avatar"`                              // 头像url
	Role          int       `gorm:"default:1;" json:"role"`                      // 用户角色：0 管理员/ 1 普通用户
	Status        int       `gorm:"default:0;" json:"status"`                    // 用户状态：0 正常/1 禁用
	CreateTime    time.Time `gorm:"autoCreateTime" json:"create_time"`           // 创建时间
	LastLoginTime time.Time `gorm:"autoUpdateTime:milli" json:"last_login_time"` // 最后登录时间
	UpdateTime    time.Time `gorm:"autoUpdateTime:milli" json:"updateTime"`      // 更新时间
	IsDeleted     int       `gorm:"default:0;" json:"is_deleted"`                // 逻辑删除（0:可用 1:不可用）
}

func (u *User) TableName() string {
	return "opbook_user"
}

func (u *User) FindByEmail(email string) (user *User) {
	global.DB.Where("is_deleted = ?", 0).Where("email = ?", email).First(&user)
	return user
}

func (u *User) FindByMobile(mobile string) (user *User) {
	global.DB.Where("is_deleted = ?", 0).Where("mobile = ?", mobile).First(&user)
	return user
}

func (u *User) CheckEmailIsExist(email string) bool {
	var count int64
	var model User
	global.DB.Model(model).Where("is_deleted = ?", 0).Where("email = ?", email).Count(&count)
	return count > 0
}

func (u *User) CheckMobileIsExist(mobile string) bool {
	var count int64
	var model User
	global.DB.Model(model).Where("is_deleted = ?", 0).Where("mobile = ?", mobile).Count(&count)
	return count > 0
}

func (u *User) Add(user *User) error {
	tx := global.DB.Create(user)
	return tx.Error
}

func (u *User) Update(user *User) error {
	return global.DB.Omit("email", "create_time", "is_deleted", "password").Where("email = ?", user.Email).Updates(user).Error
}

func (u *User) UpdatePassword(user *User) error {
	return global.DB.Select("password").Where("email = ?", user.Email).Updates(user).Error
}

func (u *User) FindByUid(uid int) (user *User) {
	global.DB.Where("is_deleted = ?", 0).Where("id = ?", uid).First(&user)
	return user
}
