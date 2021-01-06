package dto

import "novel/app/model"

type BookCreateDto struct {
	WorkDirection bool   `form:"work_direction" json:"work_direction"`          // 作品方向，0：男频，1：女频'
	CatID         int    `form:"cat_id" json:"cat_id" binding:"required"`       // 分类ID
	CatName       string `form:"cat_name" json:"cat_name"`                      // 分类名
	PicURL        string `form:"pic_url" json:"pic_url"`                        // 小说封面
	BookName      string `form:"book_name" json:"book_name" binding:"required"` // 小说名
	AuthorID      int64  `form:"author_id" json:"author_id" binding:"required"` // 作者id
	AuthorName    string `form:"author_name" json:"author_name"`                // 作者名
	BookDesc      string `form:"book_desc" json:"book_desc"`                    // 书籍描述
}

type BookCategoryRespDto struct {
	Id            int64  `json:"id"`
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
	BookId      int64 `form:"book_id"`
	BookIndexId int64 `form:"book_index_id"`
}

// 章节最新相关信息
type BookIndexAboutRespDto struct {
	BookIndexCount int64  `json:"book_index_count"` // 书的当前总的章节数
	BookContent    string `json:"book_content"`     // 章节内容
}

// 推荐书籍需要传的数据
type ListRecBookDto struct {
	CatId  int64 `form:"cat_id"`
	BookId int64 `form:"book_id"`
}

type ListRecBookRespDto struct {
	ID       int64  `json:"id"`
	PicUrl   string `json:"pic_url"`
	BookName string `json:"book_name"`
	BookDesc string `json:"book_desc"`
}

type ListCommentDto struct {
	PageDto
	BookId int64 `form:"book_id" binding:"required"`
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
