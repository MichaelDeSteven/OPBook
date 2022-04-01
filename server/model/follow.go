package model

type Follow struct {
	UserId     int  `json:"user_id"`
	FolloweeId int  `json:"followee_id"`
	FollowStat bool `json:"follow_stat"`
}

func NewFollow() *Follow {
	return &Follow{}
}
