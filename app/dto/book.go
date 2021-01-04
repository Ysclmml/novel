package dto

type BookCreateDto struct {
	WorkDirection       bool     `form:"work_direction" json:"work_direction"`     // 作品方向，0：男频，1：女频'
	CatID               int       `form:"cat_id" json:"cat_id" binding:"required"`    // 分类ID
	CatName             string    `form:"cat_name" json:"cat_name"`                // 分类名
	PicURL              string    `form:"pic_url" json:"pic_url"`                   // 小说封面
	BookName            string    `form:"book_name" json:"book_name" binding:"required"`    // 小说名
	AuthorID            int64     `form:"author_id" json:"author_id" binding:"required"`    // 作者id
	AuthorName          string    `form:"author_name" json:"author_name"` // 作者名
	BookDesc            string    `form:"book_desc" json:"book_desc"`                  // 书籍描述
}
