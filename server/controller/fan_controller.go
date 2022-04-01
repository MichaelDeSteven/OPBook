package controller

import (
	"github.com/MichaelDeSteven/OPBook/server/global"
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
	global.LOG.Sugar().Infof("%+v\n", follow)
	follow.FollowStat = socialService.FollowStatus(follow.UserId, follow.FolloweeId)
	response.OkWithData(follow, c)
}
