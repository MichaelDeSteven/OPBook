package service

import "github.com/MichaelDeSteven/OPBook/server/global"

// 计数器+1
func Increase(key string) {
	global.REDIS.Incr(key)
}
