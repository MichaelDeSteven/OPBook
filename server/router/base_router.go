package router

import (
	"net/http"
	"path"

	"github.com/MichaelDeSteven/OPBook/server/controller"
	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/middleware"
	"github.com/MichaelDeSteven/rum"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(r *rum.RouterGroup) {
	// 本地静态资源服务
	if global.CONFIG.System.OssType == "local" {
		local := global.CONFIG.Local
		dir, relativePath := local.Path, local.RelativePath
		handler := middleware.StaticHandler("filepath", relativePath, http.Dir(dir))
		urlPattern := path.Join(relativePath, "/*filepath")
		r.GET(urlPattern, handler)
	}

	searchController := controller.SearchController{}
	r.POST("/search", searchController.Result)

}
