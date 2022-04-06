package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/go-redis/redis/v8"
)

type SocialService struct{}

// 关注或取关
func (socialService *SocialService) Follow(userId, followeeId int) (follow bool) {
	global.REDIS.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		now := float64(time.Now().Unix())
		if !socialService.FollowStatus(userId, followeeId) {
			pipe.ZAdd(context.Background(), getFollowerKey(userId), &redis.Z{Score: now, Member: followeeId})
			pipe.ZAdd(context.Background(), getFolloweeKey(followeeId), &redis.Z{Score: now, Member: userId})
			follow = true
		} else {
			pipe.ZRem(context.Background(), getFollowerKey(userId), followeeId)
			pipe.ZRem(context.Background(), getFolloweeKey(followeeId), userId)
		}
		return nil
	})
	return
}

// 是否已关注
func (socialService *SocialService) FollowStatus(userId, followeeId int) bool {
	_, err := global.REDIS.ZScore(context.Background(), getFollowerKey(userId), strconv.Itoa(followeeId)).Result()
	if err != nil && err == redis.Nil {
		return false
	}
	return true
}

// 获取关注列表
func (socialService *SocialService) GetFollowees(page *model.FollowPage) (followees []model.User, totalCount uint64, err error) {
	totalCount, _ = global.REDIS.ZCard(context.Background(), getFolloweeKey(page.UserId)).Uint64()
	if page.PageIndex < 1 {
		page.PageIndex = 1
	}
	start, stop := (page.PageIndex-1)*page.PageSize, page.PageIndex+page.PageSize-2
	vals, err := global.REDIS.ZRange(context.Background(), getFolloweeKey(page.UserId), int64(start), int64(stop)).Result()
	if err != nil {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
	model := model.NewUser()
	for _, val := range vals {
		uid, _ := strconv.Atoi(val)
		followees = append(followees, *model.FindByUid(uid))
	}
	return
}

// 获取粉丝列表
func (socialService *SocialService) GetFollowers(page *model.FollowPage) (followers []model.User, totalCount uint64, err error) {
	totalCount, _ = global.REDIS.ZCard(context.Background(), getFollowerKey(page.UserId)).Uint64()
	if page.PageIndex < 1 {
		page.PageIndex = 1
	}
	start, stop := (page.PageIndex-1)*page.PageSize, page.PageIndex+page.PageSize-2
	vals, err := global.REDIS.ZRange(context.Background(), getFollowerKey(page.UserId), int64(start), int64(stop)).Result()
	if err != nil {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
	model := model.NewUser()
	for _, val := range vals {
		uid, _ := strconv.Atoi(val)
		followers = append(followers, *model.FindByUid(uid))
	}
	return
}

func getFollowerKey(userId int) string {
	return utils.FOLLOWER_SET_KEY + fmt.Sprintf("%v", userId)
}

func getFolloweeKey(followerId int) string {
	return utils.FOLLOWEE_SET_KEY + fmt.Sprintf("%v", followerId)
}

// 评论或回复
func (socialService *SocialService) CommentOrReply(comment model.Comment) {
	err := comment.AddComment()
	if err != nil {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
}

func (socialService *SocialService) DisplayComment(bookId int) (res []*model.CommentResult) {
	comment := model.NewComment()
	res = comment.GetCommentByBookId(bookId)
	ma := make(map[int]*model.CommentResult, len(res))
	for _, c := range res {
		ma[c.Id] = c
	}
	for i, r := range res {
		if res[i].CommentId != 0 {
			res[i].ReplyContent = ma[r.CommentId].Content
			res[i].ReplyNickname = ma[r.Id].Nickname
		}
	}
	return
}

// 发送私信
func (socialService *SocialService) SendMessage(message *model.Message) {
	message.ConversationId = GetConversationId(message.FromId, message.ToId)
	if err := message.Add(); err != nil {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}

}

// 获取指定会话的私信内容
func (socialService *SocialService) GetConversation(message *model.Message) (messages []*model.Message) {
	messages = message.GetConversationById(GetConversationId(message.FromId, message.ToId))
	return
}

// 获取所有与某用户有私信记录的用户列表
func (socialService *SocialService) GetConversationUserList(uid int) (list []*model.User) {
	message, user := model.NewMessage(), model.NewUser()
	mList := message.GetUserList(uid)
	uList := make([]int, len(mList))
	for i, m := range mList {
		if m.FromId == uid {
			uList[i] = m.ToId
		} else {
			uList[i] = m.FromId
		}
	}
	list = user.FindInfoByIds(uList)
	return
}

// 生成会话id
func GetConversationId(fromId, toId int) string {
	user := model.NewUser()
	if fromId < 1 || user.FindByUid(fromId) == nil {
		global.LOG.Sugar().Infof("fromId不存在: %+v\n", fromId)
		return ""
	}
	if toId < 1 || user.FindByUid(toId) == nil {
		global.LOG.Sugar().Infof("toId不存在: %+v\n", toId)
		return ""
	}
	var conversationId = ""
	if fromId < toId {
		conversationId = fmt.Sprintf("%+v_%+v", fromId, toId)
	} else {
		conversationId = fmt.Sprintf("%+v_%+v", toId, fromId)
	}
	return conversationId
}
