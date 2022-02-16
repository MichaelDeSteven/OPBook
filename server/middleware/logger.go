package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/MichaelDeSteven/rum"
)

// LogLayout 日志layout
type LogLayout struct {
	Time       time.Time
	Metadata   map[string]interface{} // 存储自定义原数据
	Path       string                 // 访问路径
	Query      string                 // 携带query
	Body       string                 // 携带body数据
	UserAgent  string                 // 代理
	Cost       string                 // 花费时间
	Source     string                 // 来源
	StatusCode int                    // 响应状态码
}

type Logger struct {
	// Filter 用户自定义过滤
	Filter func(c *rum.Context) bool
	// FilterKeyword 关键字过滤(key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess 鉴权处理
	AuthProcess func(c *rum.Context, layout *LogLayout)
	// 日志处理
	Print func(LogLayout)
	// Source 服务唯一标识
	Source string
}

func (l Logger) SetLoggerMiddleware() rum.HandlerFunc {
	return func(c *rum.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = ioutil.ReadAll(c.Request.Body)
			// 将原body塞回去
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:       time.Now(),
			Path:       path,
			Query:      query,
			UserAgent:  c.Request.UserAgent(),
			Cost:       fmt.Sprint(cost.Milliseconds()) + "ms",
			Source:     l.Source,
			StatusCode: c.StatusCode,
			Body:       string(body),
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		// 处理鉴权需要的信息
		if l.AuthProcess != nil {
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// 自行判断key/value 脱敏等
			l.FilterKeyword(&layout)
		}
		// 自行处理日志
		l.Print(layout)
	}
}

func DefaultLogger() rum.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			// 标准输出,k8s做收集
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "opbook-server",
	}.SetLoggerMiddleware()
}
