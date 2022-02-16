package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/rum"
	"go.uber.org/zap"
)

func Recovery(stack bool) rum.HandlerFunc {
	return func(c *rum.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.LOG.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					return
				}

				if stack {
					global.LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.Abort()
				c.JSON(http.StatusOK, response.Response{
					Code: http.StatusInternalServerError,
					Msg:  "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}
