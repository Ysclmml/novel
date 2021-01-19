package dto

import (
	"novel/app/model"
	"novel/app/utils/time"
	SysTime "time"
)

type BookCreateDto struct {
	WorkDirection bool   `form:"work_direction" json:"work_direction"`                 // 作品方向，0：男频，1：女频'
	CatID         int    `form:"cat_id" json:"cat_id" binding:"required"`              // 分类ID
	CatName       string `form:"cat_name" json:"cat_name"`                             // 分类名
	PicURL        string `form:"pic_url" json:"pic_url"`                               // 小说封面
	BookName      string `form:"book_name" json:"book_name string" binding:"required"` // 小说名
	AuthorID      int64  `form:"author_id" json:"author_id,string" binding:"required"` // 作者id
	AuthorName    string `form:"author_name" json:"author_name"`                       // 作者名
	BookDesc      string `form:"book_desc" json:"book_desc"`                           // 书籍描述
}

type BookCategoryRespDto struct {
	Id            int64  `json:"id,string"`
	WorkDirection bool   `json:"work_direction"` // 作品方向，0：男频，1：女频'
	Name          string `json:"name"`           // 分类名
	Sort          int8   `json:"sort"`           // 排序
}

type BookRankDto struct {
	// todo: 这里不知道为什么使用validate无效, 而使用binding才有用, 我的应该是v10版本... 而在单元测试里面用validate又有效, 迷惑..
	Type  int8 `form:"type,default=0" binding:"min=0,max=3"`
	Limit int  `form:"limit,default=30" binding:"min=2,max=50"`
}

type BookIndexAboutDto struct {
	BookId      int64 `form:"book_id,string"`
	BookIndexId int64 `form:"book_index_id,string"`
}

// 章节最新相关信息
type BookIndexAboutRespDto struct {
	BookIndexCount   int64  `json:"book_index_count"`           // 书的当前总的章节数
	BookContent      string `json:"book_content"`               // 章节内容
	BookFirstIndexId int64  `json:"book_first_index_id,string"` // 书籍首个章节id
}

// 推荐书籍需要传的数据
type ListRecBookDto struct {
	CatId  int64 `form:"cat_id,string"`
	BookId int64 `form:"book_id,string"`
}

type ListRecBookRespDto struct {
	ID       int64  `json:"id,string"`
	PicUrl   string `json:"pic_url"`
	BookName string `json:"book_name"`
	BookDesc string `json:"book_desc"`
}

type ListCommentDto struct {
	PageDto
	BookId int64 `form:"book_id" binding:"required" json:"book_id,string"`
}

type ListCommentRespDto struct {
	model.BookComment
	CreateUserName  string `json:"create_user_name"`
	CreateUserPhoto string `json:"create_user_photo"`
}

type CommentCreateDto struct {
	BookId         int64  `form:"book_id" binding:"required" json:"book_id"`
	CommentContent string `form:"comment_content" binding:"required" json:"comment_content"`
}

type IndexListRespDto struct {
	ID         int64         `json:"id,string"`
	UpdateTime time.JsonTime `json:"update_time"`
	IndexNum   int           `json:"index_num"` // 目录号
	IndexName  string        `json:"index_name"`
	IsVip      int8          `json:"is_vip"`
}

type ListIndexDto struct {
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
	BookId   int64  `form:"book_id" binding:"required" json:"book_id,string"`
	Order    string `form:"order,default=index_num"`
}

type BookSettingDto struct {
	BookId     int64   `json:"book_id,string"` // 小说ID
	Sort       int8    `json:"sort"`           // 排序号
	Type       int8    `json:"type"`
	CatID      int     `json:"cat_id"`           // 分类ID
	CatName    string  `json:"cat_name"`         // 分类名
	PicURL     string  `json:"pic_url"`          // 小说封面
	BookName   string  `json:"book_name"`        // 小说名
	AuthorID   int64   `json:"author_id,string"` // 作者id
	AuthorName string  `json:"author_name"`      // 作者名
	BookDesc   string  `json:"book_desc"`        // 书籍描述
	Score      float32 `json:"score"`            // 评分，预留字段
	BookStatus bool    `json:"book_status"`      // 书籍状态，0：连载中，1：已完结
}

// 书籍内容需要的数据
type BookContentDto struct {
	BookId  int64 `uri:"book_id" binding:"required"`
	IndexId int64 `uri:"book_index_id"  binding:"required"`
}

// 书籍分页搜索的所需的参数
type BookSP struct {
	PageDto
	WorkDirection bool         `form:"work_direction"` // 作品方向，0：男频，1：女频'
	CatId         int          `form:"cat_id"`         // 分类id
	IsVip         *bool        `form:"is_vip"`
	BookStatus    *bool        `form:"book_status"`    // 是否完结
	WordCountMin  int          `form:"word_count_min"` // 最小字数
	WordCountMax  int          `form:"word_count_max"` // 最小字数
	UpdatePeriod  int          `form:"update_period"`  // 最近更新时间, 30表示30天内更新的
	SortType      int8         `form:"sort_type"`      // 排序时间
	SortColumn    string       `form:"-"`              // 排序时间
	UpdateMinDate SysTime.Time `form:"-"`              // 最近更新时间
}
