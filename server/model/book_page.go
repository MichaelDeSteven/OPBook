package model

type BookPage struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type UserCollectPage struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
	UserId    int `json:"user_id"`
}

func NewUserCollectPage() *UserCollectPage {
	return &UserCollectPage{}
}
