package common

type PageRes struct {
	Current int `json:"current" dc:"当前页码"`
	Records any `json:"records" dc:"记录列表"`
	Total   int `json:"total" dc:"总记录数"`
	Size    int `json:"size" dc:"每页数量"`
}

type PageReq struct {
	Page int `json:"page" dc:"页码" `
	Size int `json:"size" dc:"每页数量"`
}