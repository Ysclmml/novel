package dto

import (
	"strings"
)

// GeneralListDto - General list request params
type GeneralListDto struct {
	Skip  int    `form:"skip,default=0" json:"skip"`
	Limit int    `form:"limit,default=20" json:"limit" binding:"max=10000"`
	Order string `form:"order" json:"order"`
	Q     string `form:"q" json:"q"`
}

// 分页dto
type PageDto struct {
	Page     int `form:"page,default=1" json:"page"`
	PageSize int `form:"page_size,default=10" json:"page_size" binding:"max=20"`
}

type GeneralTreeDto struct {
	Q string `form:"q" json:"q"`
}
// 通用的只传一个id的Dto
type GeneralPostDto struct {
	Id int64 `uri:"id" json:"id" binding:"required"`
}
type GeneralGetDto struct {
	Id int64 `uri:"id" json:"id" binding:"required"`
}

// TransformSearch - transform search query
func TransformSearch(qs string, mapping map[string]string) (ss map[string]string) {
	ss = make(map[string]string)
	for _, v := range strings.Split(qs, ",") {
		vs := strings.Split(v, "=")
		if _, ok := mapping[vs[0]]; ok {
			ss[mapping[vs[0]]] = vs[1]
		}
	}
	return
}
