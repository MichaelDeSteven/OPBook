package response

import (
	"net/http"

	"github.com/MichaelDeSteven/rum"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *rum.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *rum.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *rum.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *rum.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *rum.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *rum.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *rum.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *rum.Context) {
	Result(ERROR, data, message, c)
}

func FailWithError(e error, c *rum.Context) {
	Result(ERROR, map[string]interface{}{}, e.Error(), c)
}
