package model

type SearchReq struct {
	PageIndex   int    `json:"pageIndex"`
	PageSize    int    `json:"pageSize"`
	Wd          string `json:"wd"`
	IsSearchDoc bool   `json:"isSearchDoc"`
}

type SearchRes struct {
	Books       interface{} `json:"books"`
	Docs        interface{} `json:"docs"`
	SpendTime   string      `json:"spendTime"`
	Wd          string      `json:"wd"`
	TotalRows   int         `json:"totalRows"`
	IsSearchDoc bool        `json:"isSearchDoc"`
}
