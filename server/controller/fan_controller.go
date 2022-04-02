package controller

import (
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/rum"
)

type FollowController struct {
}

// 关注或取关
func (this *FollowController) Follow(c *rum.Context) {
	follow := model.NewFollow()
	c.Bind(follow)
	follow.FollowStat = socialService.Follow(follow.UserId, follow.FolloweeId)
	response.OkWithData(follow, c)
}

// 是否已关注
func (this *FollowController) FollowStatus(c *rum.Context) {
	follow := model.NewFollow()
	c.Bind(follow)
	follow.FollowStat = socialService.FollowStatus(follow.UserId, follow.FolloweeId)
	response.OkWithData(follow, c)
}

// 获取关注列表
func (this *FollowController) GetFollowees(c *rum.Context) {
	page := model.NewFollowPage()
	c.Bind(page)
	followees, totalCount, err := socialService.GetFollowees(page)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	data := make(map[string]interface{}, 2)
	data["data"], data["totalCount"] = followees, totalCount
	response.OkWithData(data, c)
}

// 获取粉丝列表
func (this *FollowController) GetFollowers(c *rum.Context) {
	page := model.NewFollowPage()
	c.Bind(page)
	followers, totalCount, err := socialService.GetFollowers(page)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	data := make(map[string]interface{}, 2)
	data["data"], data["totalCount"] = followers, totalCount
	response.OkWithData(data, c)
}
