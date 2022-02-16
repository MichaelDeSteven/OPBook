package model

type LoginInfo struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
