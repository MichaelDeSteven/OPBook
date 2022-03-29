package controller

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/OPBook/server/model/system/request"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/MichaelDeSteven/OPBook/server/utils/graphics"
	"github.com/MichaelDeSteven/OPBook/server/utils/upload"
	"github.com/MichaelDeSteven/rum"
)

func Login(c *rum.Context) {
	user := model.User{}
	c.Bind(&user)
	if user.Email == "" || user.Password == "" {
		response.FailWithMessage("用户的账号或者密码为空", c)
		return
	}
	global.LOG.Sugar().Infof("[Login] user: %+v\n", user)
	res, err := userService.Login(&user)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	tokenNext(c, res)
}

// 登录以后签发jwt
func tokenNext(c *rum.Context, user *model.User) {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.Id,
		NickName: user.Nickname,
		Email:    user.Email,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Sugar().Errorf("获取token失败! ", err)
		response.FailWithMessage("获取token失败", c)
		return
	}

	loginInfo := model.LoginInfo{
		User:  user,
		Token: token,
	}
	response.OkWithDetailed(loginInfo, "登录成功", c)
}

func Reg(c *rum.Context) {
	user := model.User{}
	c.Bind(&user)
	global.LOG.Sugar().Infof("[Reg] user: %+v\n", user)
	if user.Email == "" || user.Password == "" {
		response.FailWithMessage("用户的账号或者密码为空", c)
		return
	}
	if err := userService.Reg(&user); err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

func UpdateUserInfo(c *rum.Context) {
	user := model.User{}
	c.Bind(&user)
	email, exists := c.Get("email")
	if !exists {
		response.FailWithMessage("token无效或者过期", c)
		return
	}
	user.Email = email.(string)
	if !utils.Empty(user.Mobile) {
		if u := user.FindByMobile(user.Mobile); u != nil && email != u.Email {
			response.FailWithMessage("手机号码已存在", c)
			return
		}
	}
	if err := user.Update(&user); err != nil {
		response.FailWithMessage("更新用户信息失败", c)
		return
	}
	response.OkWithDetailed(user, "修改成功", c)
}

func Password(c *rum.Context) {
	email, exists := c.Get("email")
	if !exists {
		response.FailWithMessage("token无效或者过期", c)
		return
	}
	pw := model.Password{}
	c.Bind(&pw)
	if utils.Empty(pw.NewPassword) {
		response.FailWithMessage("新密码不能为空", c)
		return
	}
	user := &model.User{}
	user = user.FindByEmail(email.(string))
	if user.Password != pw.OldPassword {
		response.FailWithMessage("旧密码不正确！", c)
		return
	}
	if pw.OldPassword == pw.NewPassword {
		response.FailWithMessage("旧密码不能与新密码相同", c)
		return
	}
	user.Password = pw.NewPassword
	if err := user.UpdatePassword(user); err != nil {
		response.FailWithMessage("更新密码失败", c)
		return
	}
	response.OkWithDetailed(user, "修改成功", c)
}

func GetUserProfile(c *rum.Context) {
	uid := c.Param("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		response.FailWithMessage("uid应该为数字", c)
		return
	}
	user := userService.GetUserProfile(id)
	if user.Id == 0 {
		response.FailWithMessage("找不到该用户", c)
		return
	}
	response.OkWithData(user, c)
}

func UploadAvatar(c *rum.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		global.LOG.Sugar().Infof("%d", uid)
		response.FailWithMessage("不存在uid", c)
		return
	}

	form := model.AvatarForm{}
	c.Bind(&form)
	file := form.Avatar
	if file == nil {
		global.LOG.Sugar().Infof("读取文件异常!")
		response.FailWithMessage("读取文件异常", c)
		return
	}

	ext := filepath.Ext(file.Filename)
	if !strings.EqualFold(ext, ".png") &&
		!strings.EqualFold(ext, ".jpg") &&
		!strings.EqualFold(ext, ".gif") &&
		!strings.EqualFold(ext, ".jpeg") {
		global.LOG.Sugar().Infof("不支持的图片格式!")
		response.FailWithMessage("不支持的图片格式", c)
		return
	}

	x1, _ := strconv.ParseFloat(form.X, 10)
	y1, _ := strconv.ParseFloat(form.Y, 10)
	w1, _ := strconv.ParseFloat(form.Width, 10)
	h1, _ := strconv.ParseFloat(form.Height, 10)

	x := int(x1)
	y := int(y1)
	width := int(w1)
	height := int(h1)

	global.LOG.Sugar().Infof("x:%d", x)
	global.LOG.Sugar().Infof("y:%d", y)
	global.LOG.Sugar().Infof("width:%d", width)
	global.LOG.Sugar().Infof("height:%d", height)

	fileName := strconv.FormatInt(time.Now().UnixNano(), 16)

	filePath := filepath.Join("./", global.CONFIG.Local.Path, "tmp", global.CONFIG.Local.Avator, time.Now().Format("2006/01"), fileName+ext)

	path := filepath.Dir(filePath)

	os.MkdirAll(path, os.ModePerm)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		global.LOG.Sugar().Errorf("保存图片失败: %+v\n", err)
		response.FailWithMessage("保存图片失败", c)
		return
	}

	// 剪切图片
	subImg, err := graphics.ImageCopyFromFile(filePath, x, y, width, height)

	if err != nil {
		global.LOG.Sugar().Errorf("头像剪切失败: %+v\n", err)
		response.FailWithMessage("头像剪切失败", c)
	}
	os.Remove(filePath)

	// 保存剪切图片
	filePath = filepath.Join("./", global.CONFIG.Local.Path, "tmp", time.Now().Format("200601"), fileName+ext)

	err = graphics.ImageResizeSaveFile(subImg, 120, 120, filePath)
	err = graphics.SaveImage(filePath, subImg)

	if err != nil {
		global.LOG.Sugar().Errorf("保存图片失败: %+v\n", err)
		response.FailWithMessage("保存图片失败", c)
		return
	}

	oss := upload.NewOss()
	url, _, err := oss.UploadFileByPath(filePath, fileName, ext)
	if err != nil {
		global.LOG.Sugar().Errorf("保存图片失败: %+v\n", err)
		response.FailWithMessage("保存图片失败", c)
		return
	}

	// 更新用户头像链接
	user := userService.GetUserProfile(uid.(int))
	oldAvatar := user.Avatar
	user.Avatar = url
	if err := user.Update(user); err != nil {
		global.LOG.Sugar().Errorf("保存头像失败: %+v\n", err)
		response.FailWithMessage("保存头像失败", c)
		return
	}
	os.Remove(filePath)
	if err = oss.DeleteFile(oldAvatar); err != nil {
		global.LOG.Sugar().Info(err)
	}
	response.OkWithDetailed(user, "保存图片成功", c)
}
