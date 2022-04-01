package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
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

// 获取粉丝列表

func getFollowerKey(userId int) string {
	return utils.FOLLOWER_SET_KEY + fmt.Sprintf("%v", userId)
}

func getFolloweeKey(followerId int) string {
	return utils.FOLLOWEE_SET_KEY + fmt.Sprintf("%v", followerId)
}
