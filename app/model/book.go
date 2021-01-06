package model

import (
	"novel/app/utils/time"
)

// Book 小说表
type Book struct {
	BaseModel
	WorkDirection       bool           `gorm:"column:work_direction;type:tinyint(1)" json:"work_direction"`                                             // 作品方向，0：男频，1：女频'
	CatID               int            `gorm:"column:cat_id;type:int(11)" json:"cat_id"`                                                                // 分类ID
	CatName             string         `gorm:"column:cat_name;type:varchar(50)" json:"cat_name"`                                                        // 分类名
	PicURL              string         `gorm:"column:pic_url;type:varchar(200);not null" json:"pic_url"`                                                // 小说封面
	BookName            string         `gorm:"unique_index:key_uq_bookName_authorName;column:book_name;type:varchar(50);not null" json:"book_name"`     // 小说名
	AuthorID            int64          `gorm:"column:author_id;type:bigint(20)" json:"author_id"`                                                       // 作者id
	AuthorName          string         `gorm:"unique_index:key_uq_bookName_authorName;column:author_name;type:varchar(50);not null" json:"author_name"` // 作者名
	BookDesc            string         `gorm:"column:book_desc;type:varchar(2000);not null" json:"book_desc"`                                           // 书籍描述
	Score               float32        `gorm:"column:score;type:float;not null" json:"score"`                                                           // 评分，预留字段
	BookStatus          bool           `gorm:"column:book_status;type:tinyint(1);not null" json:"book_status"`                                          // 书籍状态，0：连载中，1：已完结
	VisitCount          int64          `gorm:"column:visit_count;type:bigint(20)" json:"visit_count"`                                                   // 点击量
	WordCount           int            `gorm:"column:word_count;type:int(11)" json:"word_count"`                                                        // 总字数
	CommentCount        int            `gorm:"column:comment_count;type:int(11)" json:"comment_count"`                                                  // 评论数
	YesterdayBuy        int            `gorm:"column:yesterday_buy;type:int(11)" json:"yesterday_buy"`                                                  // 昨日订阅数
	LastIndexID         int64          `gorm:"column:last_index_id;type:bigint(20)" json:"last_index_id"`                                               // 最新目录ID
	LastIndexName       string         `gorm:"column:last_index_name;type:varchar(50)" json:"last_index_name"`                                          // 最新目录名
	LastIndexUpdateTime *time.JsonTime `json:"last_index_update_time"`                                                                                  // 最新目录更新时间
	IsVip               bool           `gorm:"column:is_vip;type:tinyint(1)" json:"is_vip"`                                                             // 是否收费，1：收费，0：免费
	Status              bool           `gorm:"column:status;type:tinyint(1)" json:"status"`                                                             // 状态，0：入库，1：上架 	// 创建时间
	CrawlSourceID       int            `gorm:"column:crawl_source_id;type:int(11)" json:"-"`                                                            // 爬虫源站ID
	CrawlBookId         string         `gorm:"column:crawl_book_id;type:varchar(32)" json:"-"`                                                          // 抓取的源站小说ID
	CrawlLastTime       time.JsonTime  `json:"-"`                                                                                                       // 最后一次的抓取时间
	CrawlIsStop         bool           `gorm:"column:crawl_is_stop;type:tinyint(1)" json:"-"`                                                           // 是否已停止更新，0：未停止，1：已停止
}

func (book *Book) TableName() string {
	return "book"
}

// BookContent 小说内容表, 这章会进行分表
type BookContent struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`   // 目录章节ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`            // 小说章节内容
}

// BookIndex 小说目录表
type BookIndex struct {
	BaseModel
	BookId    int64  `gorm:"unique_index:key_uq_bookId_indexNum;index:key_bookId;column:book_id;type:bigint(20);not null" json:"book_id"`    // 小说ID
	IndexNum  int    `gorm:"unique_index:key_uq_bookId_indexNum;index:key_indexNum;column:index_num;type:int(11);not null" json:"index_num"` // 目录号
	IndexName string `gorm:"column:index_name;type:varchar(100)" json:"index_name"`                                                          // 目录名
	WordCount int    `gorm:"column:word_count;type:int(11)" json:"word_count"`                                                               // 字数
	IsVip     int8   `gorm:"column:is_vip;type:tinyint(4)" json:"is_vip"`                                                                    // 是否收费，1：收费，0：免费
	BookPrice int    `gorm:"column:book_price;type:int(3)" json:"book_price"`                                                                // 章节费用（屋币）
}

func (book *BookIndex) TableName() string {
	return "book_index"
}

// BookSetting 首页小说设置表
type BookSetting struct {
	BaseModel
	BookId       int64 `gorm:"column:book_id;type:bigint(20)" json:"book_id"`               // 小说ID
	Sort         int8  `gorm:"column:sort;type:tinyint(4)" json:"sort"`                     // 排序号
	Type         int8  `json:"type"`                                                        // 类型，0：轮播图，1：顶部小说栏设置，2：本周强推，3：热门推荐，4：精品推荐
	CreateUserId int64 `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 创建人ID
	UpdateUserId int64 `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"` // 更新人ID
}

func (book *BookSetting) TableName() string {
	return "book_setting"
}
