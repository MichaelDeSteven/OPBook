package model

type Password struct {
	OldPassword string `json:"oldPassword"` // 原始密码
	NewPassword string `json:"newPassword"` // 新密码
}
