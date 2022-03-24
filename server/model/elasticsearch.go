package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/TruthHun/gotil/util"
	"github.com/astaxie/beego/httplib"
)

// 全文搜索客户端
type ElasticSearchClient struct {
	Host           string        // host
	Index          string        // 索引
	Type           string        // type
	On             bool          // 是否启用全文搜索
	Timeout        time.Duration // 超时时间
	IsRelateSearch bool
}

// 全文搜索
type ElasticSearchData struct {
	Id       int    `json:"id"`       // 文档或书籍id
	BookId   int    `json:"book_id"`  // 书籍id。这里的book_id起到的作用是IsBooK的布尔，以及搜索书籍文档时候的过滤
	Title    string `json:"title"`    // 文档标题或书籍名称
	Keywords string `json:"keywords"` // 文档或书籍关键字
	Content  string `json:"content"`  // 文档摘要或书籍文本内容
	Vcnt     int    `json:"vcnt"`     // 浏览量
	Private  int    `json:"private"`  // 书籍或者文档是否是公开的
}

// 统计信息结构
type ElasticSearchCount struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Count int `json:"count"`
}

// 分词
type Token struct {
	EndOffset   int    `json:"end_offset"`
	Position    int    `json:"position"`
	StartOffset int    `json:"start_offset"`
	Token       string `json:"token"`
	Type        string `json:"type"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
}

// 搜索结果结构
type ElasticSearchResult struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Hits []struct {
			ID     string      `json:"_id"`
			Index  string      `json:"_index"`
			Score  interface{} `json:"_score"`
			Source struct {
				Id       int    `json:"id"`
				BookId   int    `json:"book_id"`
				Title    string `json:"title"`
				Keywords string `json:"keywords"`
				Content  string `json:"content"`
				Vcnt     int    `json:"vcnt"`
				Private  int    `json:"private"`
			} `json:"_source"`
			Type string `json:"_type"`
			Sort []int  `json:"sort"`
		} `json:"hits"`
		MaxScore interface{} `json:"max_score"`
		Total    int         `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}

// 初始化全文搜索客户端，包括检查索引是否存在，mapping设置等
func (this *ElasticSearchClient) Init() (err error) {
	if !this.On { // 未开启ElasticSearch，则不初始化
		return
	}
	// 检测是否能ping同
	if err = this.ping(); err == nil {
		// 检测索引是否存在；索引不存在，则创建索引；如果索引存在，则直接跳过初始化
		if err = this.existIndex(); err != nil {
			// 创建索引成功
			if err = this.createIndex(); err == nil {
				// 创建mapping
				js := `{
					"properties": {
						"title": {
							"type": "text",
							"analyzer": "ik_max_word",
							"search_analyzer": "ik_smart"
						},
						"keywords": {
							"type": "text",
							"analyzer": "ik_max_word",
							"search_analyzer": "ik_smart"
						},
						"description": {
							"type": "text",
							"analyzer": "ik_max_word",
							"search_analyzer": "ik_smart"
						},
						"vcnt": {
							"type": "integer"
						},
						"is_book": {
							"type": "integer"
						}
					}
				}`
				api := this.Host + this.Index + "/" + this.Type + "/_mapping"
				err = utils.HandleResponse(this.post(api).Body(js).Response())
			}
		}
	}
	return
}

// 搜索内容
// 如果书籍id大于0，则表示搜索指定的书籍的文档。否则表示搜索书籍
// 如果不指定书籍id，则只能搜索
func (this *ElasticSearchClient) Search(wd string, p, listRows int, isSearchDoc bool, bookId ...int) (result ElasticSearchResult, err error) {
	if !this.On {
		return
	}

	wd = strings.Replace(wd, "\"", " ", -1)
	wd = strings.Replace(wd, "\\", " ", -1)
	bid := 0
	if len(bookId) > 0 && bookId[0] > 0 {
		bid = bookId[0]
	}

	var queryBody string

	// 请求体
	if bid > 0 { // 搜索指定书籍的文档，不限私有和公有
		queryBody = `{"query": {"bool": {"filter": [{
				"term": {
					"book_id": {$bookId}
				}
			}],
	    "must":{"multi_match" : {
	        "query":    "%v",
            "minimum_should_match": "%v",
	        "fields": [ "title", "keywords","content" ]
	      }
	    }
		}},"from": %v,"size": %v,"_source":["id"]}`
		queryBody = strings.Replace(queryBody, "{$bookId}", strconv.Itoa(bid), 1)
	} else {
		if isSearchDoc { //搜索公开的文档
			queryBody = `{"query": {"bool": {
			"filter": [
	        {"range": {"book_id": {"gt": 0}}},
	        {"term": {"private": 0}}
	      ],"must":{
	    	"multi_match" : {
	        "query":    "%v",
			"minimum_should_match": "%v",
	        "fields": [ "title", "keywords","content" ]
	      }}}},"from": %v,"size": %v,"_source":["id"]}`
		} else { //搜索公开的书籍
			queryBody = `{"query": {"bool": {
			"filter": [
	      	{"term": {"book_id": 0}},
	          {"term": {"private": 0}}
	      ],"must":{
	    	"multi_match" : {
	        "query":    "%v",
			"minimum_should_match": "%v",
	        "fields": [ "title", "keywords","content" ]
	      }
	    }}},"from": %v, "size": %v,"_source":["id"]}`
		}
	}
	percent := "50"
	if this.IsRelateSearch {
		percent = "10"
	}
	queryBody = fmt.Sprintf(queryBody, wd, percent+"%", (p-1)*listRows, listRows)
	api := this.Host + this.Index + "/" + this.Type + "/_search"
	if resp, errResp := this.post(api).Body(queryBody).Response(); errResp != nil {
		err = errResp
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(b, &result)
	}
	return
}

//重建索引【全量】
//采用批量重建索引的方式进行
//每次操作100条数据
func (this *ElasticSearchClient) RebuildAllIndex(bookId int) {
	if !this.On || bookId == 0 {
		return
	}

	book, _ := NewBook().GetById(bookId)
	var data []ElasticSearchData
	data = append(data, ElasticSearchData{
		Id:       bookId,
		Title:    book.Name,
		Keywords: book.Label,
		Content:  book.Description,
		BookId:   0, //注意：这里必须设置为0
		Private:  book.Private,
		Vcnt:     book.ViewCount,
	})
	this.BuildIndexByBuck(data)

	// 文档内容可能比较大，每次更新10个文档
	maxPageIndex, pageSize := int(1e7), 10
	doc := NewDocument()
	for pageIndex := 1; pageIndex < maxPageIndex; pageIndex++ {
		docs, _ := doc.FindToPager(pageIndex, pageSize, bookId, "id", "book_id", "name", "release", "view_count")
		if len(docs) > 0 {
			var data []ElasticSearchData
			for _, item := range docs {
				d := ElasticSearchData{
					Id:       item.Id,
					Title:    item.Name,
					Keywords: "",
					Content:  this.html2Text(item.Release),
					BookId:   item.BookId,
					Vcnt:     item.ViewCount,
				}
				data = append(data, d)
			}
			this.BuildIndexByBuck(data)
		} else {
			pageIndex = maxPageIndex
		}
	}
}

// 通过bulk，批量创建/更新索引
func (this *ElasticSearchClient) BuildIndexByBuck(data []ElasticSearchData) (err error) {
	if !this.On {
		return
	}

	now := time.Now()
	var bodySlice []string
	if len(data) > 0 {
		var _id string
		for _, item := range data {
			// 书籍的id大于0，表示这个数据是文档的数据，否则是书籍的数据
			if item.BookId > 0 {
				_id = fmt.Sprintf("doc_%v", item.Id)
			} else {
				_id = fmt.Sprintf("book_%v", item.Id)
			}
			action := fmt.Sprintf(`{"index":{"_index":"%v","_type":"%v","_id":"%v"}}`, this.Index, this.Type, _id)
			bodySlice = append(bodySlice, action)
			bodySlice = append(bodySlice, util.InterfaceToJson(item))
		}
		api := this.Host + "_bulk"
		body := strings.Join(bodySlice, "\n") + "\n"
		err = utils.HandleResponse(this.post(api).Body(body).Response())
	}
	d := time.Since(now)
	if d > time.Duration(this.Timeout) {
		// 生成索引时长过长，休眠一小段时间
		global.LOG.Sugar().Infof("sleep second %+v\n", (time.Duration(this.Timeout/2) * time.Second).Seconds())
		time.Sleep(time.Duration(this.Timeout/2) * time.Second)
	}
	return
}

// 创建全文搜索客户端
func NewElasticSearchClient() (client *ElasticSearchClient) {
	client = &ElasticSearchClient{
		Host:    "http://localhost:9200/",
		Index:   "opbook",
		Type:    "fulltext",
		On:      true,
		Timeout: 10 * time.Second,
	}
	client.Host = strings.TrimRight(client.Host, "/") + "/"
	return
}

// 将HTML转成符合elasticsearch搜索的文本
func (this *ElasticSearchClient) html2Text(htmlStr string) string {
	var tags = []string{
		"</p>", "</div>", "</code>", "</span>", "</pre>", "</blockquote>",
		"</h1>", "</h2>", "</h3>", "</h4>", "</h5>", "</h6>", "</td>", "</th>",
		"</i>", "</b>", "</strong>", "</a>", "</li>",
	}

	for _, tag := range tags {
		htmlStr = strings.Replace(htmlStr, tag, tag+" ", -1)
	}

	htmlStr = strings.Replace(htmlStr, "\n", " ", -1)

	gq, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return htmlStr
	}
	return gq.Text()
}

//创建索引
func (this *ElasticSearchClient) BuildIndex(es ElasticSearchData) (err error) {
	if !this.On {
		return
	}

	var (
		js  []byte
		_id string
	)

	es.Content = this.html2Text(es.Content)

	if es.BookId > 0 {
		_id = fmt.Sprintf("doc_%v", es.Id)
	} else {
		_id = fmt.Sprintf("book_%v", es.Id)
	}
	api := this.Host + this.Index + "/" + this.Type + "/" + _id
	if js, err = json.Marshal(es); err == nil {
		err = utils.HandleResponse(this.post(api).Body(js).Response())
	}
	return
}

// 查询分词
func (this *ElasticSearchClient) SegWords(keyword string) string {
	if !this.On {
		return keyword
	}
	api := this.Host + this.Index + "/_analyze"
	req := this.post(api)
	keyword = strings.Replace(keyword, "\\", " ", -1)
	keyword = strings.Replace(keyword, "\"", " ", -1)
	body := `{"text":"` + keyword + `","tokenizer": "ik_max_word"}`
	req.Body(body)
	if js, err := req.String(); err == nil {
		var tokens Tokens
		json.Unmarshal([]byte(js), &tokens)
		if len(tokens.Tokens) > 0 {
			var words []string
			for _, token := range tokens.Tokens {
				words = append(words, token.Token)
			}
			return strings.Join(words, ",")
		}
	}
	return keyword
}

// 查询索引量
// @return           count           统计数据
// @return           err             错误
func (this *ElasticSearchClient) Count() (count int, err error) {
	if !this.On {
		err = errors.New("未启用ElasticSearch")
		return
	}
	api := this.Host + this.Index + "/" + this.Type + "/_count"
	if resp, errResp := this.get(api).Response(); errResp != nil {
		err = errResp
	} else {
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		body := string(b)
		if resp.StatusCode >= http.StatusMultipleChoices || resp.StatusCode < http.StatusOK {
			err = errors.New(resp.Status + "；" + body)
		} else {
			var cnt ElasticSearchCount
			if err = json.Unmarshal(b, &cnt); err == nil {
				count = cnt.Count
			}
		}
	}
	return
}

// 删除书籍索引
func (this *ElasticSearchClient) DeleteIndex(id int, isBook bool) (err error) {
	if !this.On {
		return
	}

	_id := strconv.Itoa(id)
	idStr := "doc_" + _id
	if isBook {
		idStr = "book_" + _id
	}

	// 不管是书籍id还是文档id，常规删除操作API如下：
	api := this.Host + this.Index + "/" + this.Type + "/" + idStr

	if err = utils.HandleResponse(this.delete(api).Response()); err != nil {
		global.LOG.Sugar().Errorf("api: %+v, error: %+v\n", api, err)
	}

	if isBook { // 如果是删除书籍的索引，则接下来删除书籍所对应的文档的索引。使用条件查询的方式进行删除操作
		api = this.Host + this.Index + "/_delete_by_query"
		body := fmt.Sprintf(`{"query":{"term":{ "book_id":%v}}}`, id)
		err = utils.HandleResponse(this.post(api).Body(body).Response())
		if err != nil {
			global.LOG.Sugar().Errorf("api: %+v, error: %+v\n", api, err)
		}
	}

	return
}

// 检验es服务能否连通
func (this *ElasticSearchClient) ping() error {
	return utils.HandleResponse(this.get(this.Host).Response())
}

// 查询索引是否存在
// @return			err				nil表示索引存在，否则表示不存在
func (this *ElasticSearchClient) existIndex() (err error) {
	api := this.Host + this.Index
	err = utils.HandleResponse(this.get(api).Response())
	return
}

// 创建索引
// @return           err             创建索引
func (this *ElasticSearchClient) createIndex() (err error) {
	api := this.Host + this.Index
	err = utils.HandleResponse(this.put(api).Response())
	return
}

// put请求
func (this *ElasticSearchClient) put(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Put(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

// post请求
func (this *ElasticSearchClient) post(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Post(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

// delete请求
func (this *ElasticSearchClient) delete(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Delete(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

// get请求
func (this *ElasticSearchClient) get(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Get(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}
