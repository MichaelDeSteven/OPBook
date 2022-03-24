package controller

import (
	"fmt"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/rum"
)

type SearchController struct {
}

// 搜索结果页
func (this *SearchController) Result(c *rum.Context) {
	req := model.SearchReq{}
	c.Bind(&req)
	if req.PageIndex < 1 {
		req.PageIndex = 1
	}
	req.PageSize = 10
	now := time.Now()

	client := model.NewElasticSearchClient()
	// elasticsearch 进行全文搜索
	if client.On {
		result, err := model.NewElasticSearchClient().Search(req.Wd, req.PageIndex, req.PageSize, req.IsSearchDoc)
		if err != nil {
			global.LOG.Sugar().Error(err)
			response.FailWithError(err, c)
			return
		}
		var ids []int
		for _, item := range result.Hits.Hits {
			ids = append(ids, item.Source.Id)
		}
		res := model.SearchRes{
			SpendTime:   fmt.Sprintf("%.3f", time.Since(now).Seconds()),
			Wd:          req.Wd,
			TotalRows:   result.Hits.Total,
			IsSearchDoc: req.IsSearchDoc,
		}
		if req.IsSearchDoc {
			docs := model.NewDocResult().GetDocsById(ids)
			res.Docs = docs
		} else {
			books, _ := model.NewBook().GetBooksById(ids)
			res.Books = books
		}
		response.OkWithData(res, c)
	} else {
		global.LOG.Sugar().Infof("未开启全文搜索\n")
		response.FailWithMessage("未开启全文搜索\n", c)
	}
}
