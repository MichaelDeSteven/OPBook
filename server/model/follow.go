package model

type Follow struct {
	UserId     int  `json:"user_id"`
	FolloweeId int  `json:"followee_id"`
	FollowStat bool `json:"follow_stat"`
}

type FollowPage struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
	UserId    int `json:"user_id"`
}

func NewFollow() *Follow {
	return &Follow{}
}

func NewFollowPage() *FollowPage {
	return &FollowPage{}
}
