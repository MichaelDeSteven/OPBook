package middleware

import (
	"strconv"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/MichaelDeSteven/rum"
)

var HEADER = "x-token"

func JWTAuth() rum.HandlerFunc {
	return func(c *rum.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get(HEADER)
		if token == "" {
			response.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}
		// global.LOG.Info("token:", zap.Any("token", token))
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithMessage("授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(rum.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.SetHeader(HEADER, newToken)
			c.SetHeader("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Set("uid", claims.ID)
		c.Set("email", claims.Email)
		c.Next()
	}
}
