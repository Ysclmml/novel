package database

import (
	"time"
)

// Author 作者表
type Author struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`     // 主键
	UserID        int64     `gorm:"column:user_id;type:bigint(20)" json:"user_id"`               // 用户ID
	InviteCode    string    `gorm:"column:invite_code;type:varchar(20)" json:"invite_code"`      // 邀请码
	PenName       string    `gorm:"column:pen_name;type:varchar(20)" json:"pen_name"`            // 笔名
	TelPhone      string    `gorm:"column:tel_phone;type:varchar(20)" json:"tel_phone"`          // 手机号码
	ChatAccount   string    `gorm:"column:chat_account;type:varchar(50)" json:"chat_account"`    // QQ或微信账号
	Email         string    `gorm:"column:email;type:varchar(50)" json:"email"`                  // 电子邮箱
	WorkDirection int8      `gorm:"column:work_direction;type:tinyint(4)" json:"work_direction"` // 作品方向，0：男频，1：女频
	Status        int8      `gorm:"column:status;type:tinyint(4)" json:"status"`                 // 0：正常，1：封禁
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 创建时间
}

// AuthorCode 作家邀请码表
type AuthorCode struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`        // 主键
	InviteCode   string    `gorm:"unique;column:invite_code;type:varchar(100)" json:"invite_code"` // 邀请码
	ValidityTime time.Time `gorm:"column:validity_time;type:datetime" json:"validity_time"`        // 有效时间
	IsUse        bool      `gorm:"column:is_use;type:tinyint(1)" json:"is_use"`                    // 是否使用过，0：未使用，1:使用过
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 创建时间
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"`    // 创建人ID
}

// AuthorIncome 稿费收入统计表
type AuthorIncome struct {
	ID             int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                  // 主键
	UserID         int64     `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`                   // 用户ID
	AuthorID       int64     `gorm:"column:author_id;type:bigint(20);not null" json:"author_id"`               // 作家ID
	BookID         int64     `gorm:"column:book_id;type:bigint(20);not null" json:"book_id"`                   // 作品ID
	IncomeMonth    time.Time `gorm:"column:income_month;type:date;not null" json:"income_month"`               // 收入月份
	PreTaxIncome   int64     `gorm:"column:pre_tax_income;type:bigint(20);not null" json:"pre_tax_income"`     // 税前收入（分）
	AfterTaxIncome int64     `gorm:"column:after_tax_income;type:bigint(20);not null" json:"after_tax_income"` // 税后收入（分）
	PayStatus      bool      `gorm:"column:pay_status;type:tinyint(1);not null" json:"pay_status"`             // 支付状态，0：待支付，1：已支付
	ConfirmStatus  bool      `gorm:"column:confirm_status;type:tinyint(1);not null" json:"confirm_status"`     // 稿费确认状态，0：待确认，1：已确认
	Detail         string    `gorm:"column:detail;type:varchar(255)" json:"detail"`                            // 详情
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
}

// AuthorIncomeDetail 稿费收入明细统计表
type AuthorIncomeDetail struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`           // 主键
	UserID        int64     `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`            // 用户ID
	AuthorID      int64     `gorm:"column:author_id;type:bigint(20);not null" json:"author_id"`        // 作家ID
	BookID        int64     `gorm:"column:book_id;type:bigint(20);not null" json:"book_id"`            // 作品ID,0表示全部作品
	IncomeDate    time.Time `gorm:"column:income_date;type:date;not null" json:"income_date"`          // 收入日期
	IncomeAccount int       `gorm:"column:income_account;type:int(11);not null" json:"income_account"` // 订阅总额
	IncomeCount   int       `gorm:"column:income_count;type:int(11);not null" json:"income_count"`     // 订阅次数
	IncomeNumber  int       `gorm:"column:income_number;type:int(11);not null" json:"income_number"`   // 订阅人数
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
}

// Book 小说表
type Book struct {
	ID                  int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                                 // 主键
	WorkDirection       bool      `gorm:"column:work_direction;type:tinyint(1)" json:"work_direction"`                                             // 作品方向，0：男频，1：女频'
	CatID               int       `gorm:"column:cat_id;type:int(11)" json:"cat_id"`                                                                // 分类ID
	CatName             string    `gorm:"column:cat_name;type:varchar(50)" json:"cat_name"`                                                        // 分类名
	PicURL              string    `gorm:"column:pic_url;type:varchar(200);not null" json:"pic_url"`                                                // 小说封面
	BookName            string    `gorm:"unique_index:key_uq_bookName_authorName;column:book_name;type:varchar(50);not null" json:"book_name"`     // 小说名
	AuthorID            int64     `gorm:"column:author_id;type:bigint(20)" json:"author_id"`                                                       // 作者id
	AuthorName          string    `gorm:"unique_index:key_uq_bookName_authorName;column:author_name;type:varchar(50);not null" json:"author_name"` // 作者名
	BookDesc            string    `gorm:"column:book_desc;type:varchar(2000);not null" json:"book_desc"`                                           // 书籍描述
	Score               float32   `gorm:"column:score;type:float;not null" json:"score"`                                                           // 评分，预留字段
	BookStatus          bool      `gorm:"column:book_status;type:tinyint(1);not null" json:"book_status"`                                          // 书籍状态，0：连载中，1：已完结
	VisitCount          int64     `gorm:"column:visit_count;type:bigint(20)" json:"visit_count"`                                                   // 点击量
	WordCount           int       `gorm:"column:word_count;type:int(11)" json:"word_count"`                                                        // 总字数
	CommentCount        int       `gorm:"column:comment_count;type:int(11)" json:"comment_count"`                                                  // 评论数
	YesterdayBuy        int       `gorm:"column:yesterday_buy;type:int(11)" json:"yesterday_buy"`                                                  // 昨日订阅数
	LastIndexID         int64     `gorm:"column:last_index_id;type:bigint(20)" json:"last_index_id"`                                               // 最新目录ID
	LastIndexName       string    `gorm:"column:last_index_name;type:varchar(50)" json:"last_index_name"`                                          // 最新目录名
	LastIndexUpdateTime time.Time `gorm:"index:key_lastIndexUpdateTime;column:last_index_update_time;type:datetime" json:"last_index_update_time"` // 最新目录更新时间
	IsVip               bool      `gorm:"column:is_vip;type:tinyint(1)" json:"is_vip"`                                                             // 是否收费，1：收费，0：免费
	Status              bool      `gorm:"column:status;type:tinyint(1)" json:"status"`                                                             // 状态，0：入库，1：上架
	UpdateTime          time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`                                            // 更新时间
	CreateTime          time.Time `gorm:"index:key_createTime;column:create_time;type:datetime" json:"create_time"`                                // 创建时间
	CrawlSourceID       int       `gorm:"column:crawl_source_id;type:int(11)" json:"crawl_source_id"`                                              // 爬虫源站ID
	CrawlBookID         string    `gorm:"column:crawl_book_id;type:varchar(32)" json:"crawl_book_id"`                                              // 抓取的源站小说ID
	CrawlLastTime       time.Time `gorm:"column:crawl_last_time;type:datetime" json:"crawl_last_time"`                                             // 最后一次的抓取时间
	CrawlIsStop         bool      `gorm:"column:crawl_is_stop;type:tinyint(1)" json:"crawl_is_stop"`                                               // 是否已停止更新，0：未停止，1：已停止
}

// BookAuthor 作者表
type BookAuthor struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`     // 主键
	InviteCode    string    `gorm:"column:invite_code;type:varchar(20)" json:"invite_code"`      // 邀请码
	PenName       string    `gorm:"column:pen_name;type:varchar(20)" json:"pen_name"`            // 笔名
	TelPhone      string    `gorm:"column:tel_phone;type:varchar(20)" json:"tel_phone"`          // 手机号码
	ChatAccount   string    `gorm:"column:chat_account;type:varchar(50)" json:"chat_account"`    // QQ或微信账号
	Email         string    `gorm:"column:email;type:varchar(50)" json:"email"`                  // 电子邮箱
	WorkDirection int8      `gorm:"column:work_direction;type:tinyint(4)" json:"work_direction"` // 作品方向，0：男频，1：女频
	Status        int8      `gorm:"column:status;type:tinyint(4)" json:"status"`                 // 0：待审核，1：审核通过，正常，2：审核不通过
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 创建时间
	CreateUserID  int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 申请人ID
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`         // 更新时间
	UpdateUserID  int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"` // 更新人ID
}

// BookCategory 小说类别表
type BookCategory struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`        // 主键
	WorkDirection bool      `gorm:"column:work_direction;type:tinyint(1)" json:"work_direction"` // 作品方向，0：男频，1：女频'
	Name          string    `gorm:"column:name;type:varchar(20);not null" json:"name"`           // 分类名
	Sort          int8      `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`            // 排序
	CreateUserID  int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateUserID  int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

// BookComment 小说评论表
type BookComment struct {
	ID             int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                       // 主键
	BookID         int64     `gorm:"unique_index:key_uq_bookid_userid;column:book_id;type:bigint(20)" json:"book_id"`               // 小说ID
	CommentContent string    `gorm:"column:comment_content;type:varchar(512)" json:"comment_content"`                               // 评价内容
	ReplyCount     int       `gorm:"column:reply_count;type:int(11)" json:"reply_count"`                                            // 回复数量
	AuditStatus    bool      `gorm:"column:audit_status;type:tinyint(1)" json:"audit_status"`                                       // 审核状态，0：待审核，1：审核通过，2：审核不通过
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`                                           // 评价时间
	CreateUserID   int64     `gorm:"unique_index:key_uq_bookid_userid;column:create_user_id;type:bigint(20)" json:"create_user_id"` // 评价人
}

// BookCommentReply 小说评论回复表
type BookCommentReply struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`     // 主键
	CommentID    int64     `gorm:"column:comment_id;type:bigint(20)" json:"comment_id"`         // 评论ID
	ReplyContent string    `gorm:"column:reply_content;type:varchar(512)" json:"reply_content"` // 回复内容
	AuditStatus  bool      `gorm:"column:audit_status;type:tinyint(1)" json:"audit_status"`     // 审核状态，0：待审核，1：审核通过，2：审核不通过
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 回复用户ID
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 回复时间
}

// BookContent 小说内容表
type BookContent struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent0 小说内容表
type BookContent0 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent1 小说内容表
type BookContent1 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent2 小说内容表
type BookContent2 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent3 小说内容表
type BookContent3 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent4 小说内容表
type BookContent4 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent5 小说内容表
type BookContent5 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent6 小说内容表
type BookContent6 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent7 小说内容表
type BookContent7 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent8 小说内容表
type BookContent8 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookContent9 小说内容表
type BookContent9 struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键
	IndexID int64  `gorm:"unique;column:index_id;type:bigint(20)" json:"index_id"`  // 目录ID
	Content string `gorm:"column:content;type:mediumtext" json:"content"`           // 小说章节内容
}

// BookIndex 小说目录表
type BookIndex struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                                        // 主键
	BookID     int64     `gorm:"unique_index:key_uq_bookId_indexNum;index:key_bookId;column:book_id;type:bigint(20);not null" json:"book_id"`    // 小说ID
	IndexNum   int       `gorm:"unique_index:key_uq_bookId_indexNum;index:key_indexNum;column:index_num;type:int(11);not null" json:"index_num"` // 目录号
	IndexName  string    `gorm:"column:index_name;type:varchar(100)" json:"index_name"`                                                          // 目录名
	WordCount  int       `gorm:"column:word_count;type:int(11)" json:"word_count"`                                                               // 字数
	IsVip      int8      `gorm:"column:is_vip;type:tinyint(4)" json:"is_vip"`                                                                    // 是否收费，1：收费，0：免费
	BookPrice  int       `gorm:"column:book_price;type:int(3)" json:"book_price"`                                                                // 章节费用（屋币）
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

// BookScreenBullet 小说弹幕表
type BookScreenBullet struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                          // 主键
	ContentID    int64     `gorm:"index:key_contentId;column:content_id;type:bigint(20);not null" json:"content_id"` // 小说内容ID
	ScreenBullet string    `gorm:"column:screen_bullet;type:varchar(512);not null" json:"screen_bullet"`             // 小说弹幕内容
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`                     // 创建时间
}

// BookSetting 首页小说设置表
type BookSetting struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	BookID       int64     `gorm:"column:book_id;type:bigint(20)" json:"book_id"`               // 小说ID
	Sort         int8      `gorm:"column:sort;type:tinyint(4)" json:"sort"`                     // 排序号
	Type         bool      `gorm:"column:type;type:tinyint(1)" json:"type"`                     // 类型，0：轮播图，1：顶部小说栏设置，2：本周强推，3：热门推荐，4：精品推荐
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 创建时间
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 创建人ID
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`         // 更新时间
	UpdateUserID int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"` // 更新人ID
}

// CrawlBatchTask 批量抓取任务表
type CrawlBatchTask struct {
	ID                int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`            // 主键
	SourceID          int       `gorm:"column:source_id;type:int(11)" json:"source_id"`                     // 爬虫源ID
	CrawlCountSuccess int       `gorm:"column:crawl_count_success;type:int(11)" json:"crawl_count_success"` // 成功抓取数量
	CrawlCountTarget  int       `gorm:"column:crawl_count_target;type:int(11)" json:"crawl_count_target"`   // 目标抓取数量
	TaskStatus        bool      `gorm:"column:task_status;type:tinyint(1)" json:"task_status"`              // 任务状态，1：正在运行，0已停止
	StartTime         time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`                  // 任务开始时间
	EndTime           time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                      // 任务结束时间
}

// CrawlSingleTask 抓取单本小说任务表
type CrawlSingleTask struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`       // 主键
	SourceID     int       `gorm:"column:source_id;type:int(11)" json:"source_id"`                // 爬虫源ID
	SourceName   string    `gorm:"column:source_name;type:varchar(50)" json:"source_name"`        // 爬虫源名
	SourceBookID string    `gorm:"column:source_book_id;type:varchar(255)" json:"source_book_id"` // 源站小说ID
	CatID        int       `gorm:"column:cat_id;type:int(11)" json:"cat_id"`                      // 分类ID
	BookName     string    `gorm:"column:book_name;type:varchar(50)" json:"book_name"`            // 爬取的小说名
	AuthorName   string    `gorm:"column:author_name;type:varchar(50)" json:"author_name"`        // 爬取的小说作者名
	TaskStatus   bool      `gorm:"column:task_status;type:tinyint(1)" json:"task_status"`         // 任务状态，0：失败，1：成功，2；未执行
	ExcCount     int8      `gorm:"column:exc_count;type:tinyint(2)" json:"exc_count"`             // 已经执行次数，最多执行5次
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`           // 创建时间
}

// CrawlSource 爬虫源表
type CrawlSource struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`      // 主键
	SourceName   string    `gorm:"column:source_name;type:varchar(50)" json:"source_name"`    // 源站名
	CrawlRule    string    `gorm:"column:crawl_rule;type:text" json:"crawl_rule"`             // 爬取规则（json串）
	SourceStatus bool      `gorm:"column:source_status;type:tinyint(1)" json:"source_status"` // 爬虫源状态，0：关闭，1：开启
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`       // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`       // 更新时间
}

// FriendLink [...]
type FriendLink struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`        // 主键
	LinkName     string    `gorm:"column:link_name;type:varchar(50);not null" json:"link_name"` // 链接名
	LinkURL      string    `gorm:"column:link_url;type:varchar(100);not null" json:"link_url"`  // 链接url
	Sort         int8      `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`            // 排序号
	IsOpen       bool      `gorm:"column:is_open;type:tinyint(1);not null" json:"is_open"`      // 是否开启，0：不开启，1：开启
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 创建人id
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 创建时间
	UpdateUserID int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"` // 更新者用户id
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`         // 更新时间
}

// News 新闻表
type News struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`     // 主键
	CatID        int       `gorm:"column:cat_id;type:int(11)" json:"cat_id"`                    // 类别ID
	CatName      string    `gorm:"column:cat_name;type:varchar(50)" json:"cat_name"`            // 分类名
	SourceName   string    `gorm:"column:source_name;type:varchar(50)" json:"source_name"`      // 来源
	Title        string    `gorm:"column:title;type:varchar(100)" json:"title"`                 // 标题
	Content      string    `gorm:"column:content;type:text" json:"content"`                     // 内容
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 发布时间
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 发布人ID
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`         // 更新时间
	UpdateUserID int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"` // 更新人ID
}

// NewsCategory 新闻类别表
type NewsCategory struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"` // 主键
	Name         string    `gorm:"column:name;type:varchar(20);not null" json:"name"`    // 分类名
	Sort         int8      `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`     // 排序
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateUserID int64     `gorm:"column:update_user_id;type:bigint(20)" json:"update_user_id"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

// OrderPay 充值订单
type OrderPay struct {
	ID          int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`          // 主键
	OutTradeNo  int64     `gorm:"column:out_trade_no;type:bigint(20);not null" json:"out_trade_no"` // 商户订单号
	TradeNo     string    `gorm:"column:trade_no;type:varchar(64)" json:"trade_no"`                 // 支付宝/微信交易号
	PayChannel  bool      `gorm:"column:pay_channel;type:tinyint(1);not null" json:"pay_channel"`   // 支付渠道，1：支付宝，2：微信
	TotalAmount int       `gorm:"column:total_amount;type:int(11);not null" json:"total_amount"`    // 交易金额(单位元)
	UserID      int64     `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`           // 支付用户ID
	PayStatus   bool      `gorm:"column:pay_status;type:tinyint(1)" json:"pay_status"`              // 支付状态：0：支付失败，1：支付成功，2：待支付
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`              // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`              // 更新时间
}

// SysDataPerm 数据权限管理
type SysDataPerm struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	Name          string    `gorm:"column:name;type:varchar(50)" json:"name"`                       // 权限名称
	TableName     string    `gorm:"column:table_name;type:varchar(50)" json:"table_name"`           // 数据表名称
	ModuleName    string    `gorm:"column:module_name;type:varchar(50)" json:"module_name"`         // 所属模块
	CrlAttrName   string    `gorm:"column:crl_attr_name;type:varchar(50)" json:"crl_attr_name"`     // 用户权限控制属性名
	CrlColumnName string    `gorm:"column:crl_column_name;type:varchar(50)" json:"crl_column_name"` // 数据表权限控制列名
	PermCode      string    `gorm:"column:perm_code;type:varchar(50)" json:"perm_code"`             // 权限code，all_开头表示查看所有数据的权限，sup_开头表示查看下级数据的权限，own_开头表示查看本级数据的权限
	OrderNum      int       `gorm:"column:order_num;type:int(11)" json:"order_num"`                 // 排序
	GmtCreate     time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`              // 创建时间
	GmtModified   time.Time `gorm:"column:gmt_modified;type:datetime" json:"gmt_modified"`          // 修改时间
}

// SysDept 部门管理
type SysDept struct {
	DeptID   int64  `gorm:"primary_key;column:dept_id;type:bigint(20);not null" json:"-"`
	ParentID int64  `gorm:"column:parent_id;type:bigint(20)" json:"parent_id"` // 上级部门ID，一级部门为0
	Name     string `gorm:"column:name;type:varchar(50)" json:"name"`          // 部门名称
	OrderNum int    `gorm:"column:order_num;type:int(11)" json:"order_num"`    // 排序
	DelFlag  int8   `gorm:"column:del_flag;type:tinyint(4)" json:"del_flag"`   // 是否删除  -1：已删除  0：正常
}

// SysDict 字典表
type SysDict struct {
	ID          int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`              // 编号
	Name        string    `gorm:"index:sys_dict_label;column:name;type:varchar(100)" json:"name"`       // 标签名
	Value       string    `gorm:"index:sys_dict_value;column:value;type:varchar(100)" json:"value"`     // 数据值
	Type        string    `gorm:"column:type;type:varchar(100)" json:"type"`                            // 类型
	Description string    `gorm:"column:description;type:varchar(100)" json:"description"`              // 描述
	Sort        float64   `gorm:"column:sort;type:decimal(10,0)" json:"sort"`                           // 排序（升序）
	ParentID    int64     `gorm:"column:parent_id;type:bigint(20)" json:"parent_id"`                    // 父级编号
	CreateBy    int       `gorm:"column:create_by;type:int(11)" json:"create_by"`                       // 创建者
	CreateDate  time.Time `gorm:"column:create_date;type:datetime" json:"create_date"`                  // 创建时间
	UpdateBy    int64     `gorm:"column:update_by;type:bigint(20)" json:"update_by"`                    // 更新者
	UpdateDate  time.Time `gorm:"column:update_date;type:datetime" json:"update_date"`                  // 更新时间
	Remarks     string    `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                      // 备注信息
	DelFlag     string    `gorm:"index:sys_dict_del_flag;column:del_flag;type:char(1)" json:"del_flag"` // 删除标记
}

// SysFile 文件上传
type SysFile struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	Type       int       `gorm:"column:type;type:int(11)" json:"type"`                // 文件类型
	URL        string    `gorm:"column:url;type:varchar(200)" json:"url"`             // URL地址
	CreateDate time.Time `gorm:"column:create_date;type:datetime" json:"create_date"` // 创建时间
}

// SysGenColumns [...]
type SysGenColumns struct {
	ID            int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`        // 主键
	TableName     string `gorm:"column:table_name;type:varchar(64)" json:"table_name"`           // 表名
	ColumnName    string `gorm:"column:column_name;type:varchar(64)" json:"column_name"`         // 列名
	ColumnType    string `gorm:"column:column_type;type:varchar(64)" json:"column_type"`         // 列类型
	JavaType      string `gorm:"column:java_type;type:varchar(64)" json:"java_type"`             // 映射java类型
	ColumnComment string `gorm:"column:column_comment;type:varchar(1024)" json:"column_comment"` // 列注释
	ColumnSort    int8   `gorm:"column:column_sort;type:tinyint(4)" json:"column_sort"`          // 列排序（升序）
	ColumnLabel   string `gorm:"column:column_label;type:varchar(64)" json:"column_label"`       // 鍒楁爣绛惧悕
	PageType      int8   `gorm:"column:page_type;type:tinyint(4)" json:"page_type"`              // 页面显示类型：1、文本框 2、下拉框 3、数值4、日期 5、文本域6、富文本 7、上传图片【单文件】 8、上传图片【多文件】9、上传文件【单文件】 10、上传文件【多文件】11、隐藏域 12、不显示
	IsRequired    bool   `gorm:"column:is_required;type:tinyint(1)" json:"is_required"`          // 是否必填
	DictType      string `gorm:"column:dict_type;type:varchar(100)" json:"dict_type"`            // 页面显示为下拉时使用，字典类型从字典表中取出
}

// SysGenTable 代码生成表
type SysGenTable struct {
	ID                 int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                  // 主键
	TableName          string    `gorm:"column:table_name;type:varchar(64);not null" json:"table_name"`            // 表名
	ClassName          string    `gorm:"column:class_name;type:varchar(100);not null" json:"class_name"`           // 实体类名称
	Comments           string    `gorm:"column:comments;type:varchar(500);not null" json:"comments"`               // 表说明
	Category           bool      `gorm:"column:category;type:tinyint(1);not null" json:"category"`                 // 分类：0：数据表，1：树表
	PackageName        string    `gorm:"column:package_name;type:varchar(500)" json:"package_name"`                // 生成包路径
	ModuleName         string    `gorm:"column:module_name;type:varchar(30)" json:"module_name"`                   // 生成模块名
	SubModuleName      string    `gorm:"column:sub_module_name;type:varchar(30)" json:"sub_module_name"`           // 生成子模块名
	FunctionName       string    `gorm:"column:function_name;type:varchar(200)" json:"function_name"`              // 生成功能名，用于类描述
	FunctionNameSimple string    `gorm:"column:function_name_simple;type:varchar(50)" json:"function_name_simple"` // 生成功能名（简写），用于功能提示，如“保存xx成功”
	Author             string    `gorm:"column:author;type:varchar(50)" json:"author"`                             // 生成功能作者
	SrcDir             string    `gorm:"column:src_dir;type:varchar(1000)" json:"src_dir"`                         // src目录
	Options            string    `gorm:"column:options;type:varchar(1000)" json:"options"`                         // 其它生成选项
	CreateBy           int64     `gorm:"column:create_by;type:bigint(20);not null" json:"create_by"`               // 创建者
	CreateDate         time.Time `gorm:"column:create_date;type:datetime;not null" json:"create_date"`             // 创建时间
	UpdateBy           int64     `gorm:"column:update_by;type:bigint(20);not null" json:"update_by"`               // 更新者
	UpdateDate         time.Time `gorm:"column:update_date;type:datetime;not null" json:"update_date"`             // 更新时间
	Remarks            string    `gorm:"column:remarks;type:varchar(500)" json:"remarks"`                          // 备注信息
}

// SysGenTableColumn 代码生成表列
type SysGenTableColumn struct {
	ID          int64   `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                // 主键
	TableID     int64   `gorm:"index:idx_gen_table_column_tn;column:table_id;type:bigint(20);not null" json:"table_id"` // 表id
	ColumnName  string  `gorm:"column:column_name;type:varchar(64);not null" json:"column_name"`                        // 列名
	ColumnSort  float64 `gorm:"column:column_sort;type:decimal(10,0)" json:"column_sort"`                               // 列排序（升序）
	ColumnType  string  `gorm:"column:column_type;type:varchar(100);not null" json:"column_type"`                       // 类型
	ColumnLabel string  `gorm:"column:column_label;type:varchar(50)" json:"column_label"`                               // 列标签名
	Comments    string  `gorm:"column:comments;type:varchar(500);not null" json:"comments"`                             // 列备注说明
	AttrName    string  `gorm:"column:attr_name;type:varchar(200);not null" json:"attr_name"`                           // 类的属性名
	AttrType    string  `gorm:"column:attr_type;type:varchar(200);not null" json:"attr_type"`                           // 类的属性类型
	IsPk        string  `gorm:"column:is_pk;type:char(1)" json:"is_pk"`                                                 // 是否主键
	IsNull      string  `gorm:"column:is_null;type:char(1)" json:"is_null"`                                             // 是否可为空
	IsInsert    string  `gorm:"column:is_insert;type:char(1)" json:"is_insert"`                                         // 是否插入字段
	IsUpdate    string  `gorm:"column:is_update;type:char(1)" json:"is_update"`                                         // 是否更新字段
	IsList      string  `gorm:"column:is_list;type:char(1)" json:"is_list"`                                             // 是否列表字段
	IsQuery     string  `gorm:"column:is_query;type:char(1)" json:"is_query"`                                           // 是否查询字段
	QueryType   string  `gorm:"column:query_type;type:varchar(200)" json:"query_type"`                                  // 查询方式
	IsEdit      string  `gorm:"column:is_edit;type:char(1)" json:"is_edit"`                                             // 是否编辑字段
	ShowType    string  `gorm:"column:show_type;type:varchar(200)" json:"show_type"`                                    // 表单类型
	Options     string  `gorm:"column:options;type:varchar(1000)" json:"options"`                                       // 其它生成选项
}

// SysLog 系统日志
type SysLog struct {
	ID        int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	UserID    int64     `gorm:"column:user_id;type:bigint(20)" json:"user_id"`      // 用户id
	Username  string    `gorm:"column:username;type:varchar(50)" json:"username"`   // 用户名
	Operation string    `gorm:"column:operation;type:varchar(50)" json:"operation"` // 用户操作
	Time      int       `gorm:"column:time;type:int(11)" json:"time"`               // 响应时间
	Method    string    `gorm:"column:method;type:varchar(200)" json:"method"`      // 请求方法
	Params    string    `gorm:"column:params;type:varchar(5000)" json:"params"`     // 请求参数
	IP        string    `gorm:"column:ip;type:varchar(64)" json:"ip"`               // IP地址
	GmtCreate time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`  // 创建时间
}

// SysMenu 菜单管理
type SysMenu struct {
	MenuID      int64     `gorm:"primary_key;column:menu_id;type:bigint(20);not null" json:"-"`
	ParentID    int64     `gorm:"column:parent_id;type:bigint(20)" json:"parent_id"`     // 父菜单ID，一级菜单为0
	Name        string    `gorm:"column:name;type:varchar(50)" json:"name"`              // 菜单名称
	URL         string    `gorm:"column:url;type:varchar(200)" json:"url"`               // 菜单URL
	Perms       string    `gorm:"column:perms;type:varchar(500)" json:"perms"`           // 授权(多个用逗号分隔，如：user:list,user:create)
	Type        int       `gorm:"column:type;type:int(11)" json:"type"`                  // 类型   0：目录   1：菜单   2：按钮
	Icon        string    `gorm:"column:icon;type:varchar(50)" json:"icon"`              // 菜单图标
	OrderNum    int       `gorm:"column:order_num;type:int(11)" json:"order_num"`        // 排序
	GmtCreate   time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`     // 创建时间
	GmtModified time.Time `gorm:"column:gmt_modified;type:datetime" json:"gmt_modified"` // 修改时间
}

// SysRole 角色
type SysRole struct {
	RoleID       int64     `gorm:"primary_key;column:role_id;type:bigint(20);not null" json:"-"`
	RoleName     string    `gorm:"column:role_name;type:varchar(100)" json:"role_name"`         // 角色名称
	RoleSign     string    `gorm:"column:role_sign;type:varchar(100)" json:"role_sign"`         // 角色标识
	Remark       string    `gorm:"column:remark;type:varchar(100)" json:"remark"`               // 备注
	UserIDCreate int64     `gorm:"column:user_id_create;type:bigint(20)" json:"user_id_create"` // 创建用户id
	GmtCreate    time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`           // 创建时间
	GmtModified  time.Time `gorm:"column:gmt_modified;type:datetime" json:"gmt_modified"`       // 创建时间
}

// SysRoleDataPerm 角色与数据权限对应关系
type SysRoleDataPerm struct {
	ID     int64 `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	RoleID int64 `gorm:"column:role_id;type:bigint(20)" json:"role_id"` // 角色ID
	PermID int64 `gorm:"column:perm_id;type:bigint(20)" json:"perm_id"` // 权限ID
}

// SysRoleMenu 角色与菜单对应关系
type SysRoleMenu struct {
	ID     int64 `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	RoleID int64 `gorm:"column:role_id;type:bigint(20)" json:"role_id"` // 角色ID
	MenuID int64 `gorm:"column:menu_id;type:bigint(20)" json:"menu_id"` // 菜单ID
}

// SysUser [...]
type SysUser struct {
	UserID       int64     `gorm:"primary_key;column:user_id;type:bigint(20);not null" json:"-"`
	Username     string    `gorm:"column:username;type:varchar(50)" json:"username"` // 用户名
	Name         string    `gorm:"column:name;type:varchar(100)" json:"name"`
	Password     string    `gorm:"column:password;type:varchar(50)" json:"password"` // 密码
	DeptID       int64     `gorm:"column:dept_id;type:bigint(20)" json:"dept_id"`
	Email        string    `gorm:"column:email;type:varchar(100)" json:"email"`                 // 邮箱
	Mobile       string    `gorm:"column:mobile;type:varchar(100)" json:"mobile"`               // 手机号
	Status       int8      `gorm:"column:status;type:tinyint(4)" json:"status"`                 // 状态 0:禁用，1:正常
	UserIDCreate int64     `gorm:"column:user_id_create;type:bigint(20)" json:"user_id_create"` // 创建用户id
	GmtCreate    time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`           // 创建时间
	GmtModified  time.Time `gorm:"column:gmt_modified;type:datetime" json:"gmt_modified"`       // 修改时间
	Sex          int64     `gorm:"column:sex;type:bigint(20)" json:"sex"`                       // 性别
	Birth        time.Time `gorm:"column:birth;type:datetime" json:"birth"`                     // 出身日期
	PicID        int64     `gorm:"column:pic_id;type:bigint(20)" json:"pic_id"`
	LiveAddress  string    `gorm:"column:live_address;type:varchar(500)" json:"live_address"` // 现居住地
	Hobby        string    `gorm:"column:hobby;type:varchar(255)" json:"hobby"`               // 爱好
	Province     string    `gorm:"column:province;type:varchar(255)" json:"province"`         // 省份
	City         string    `gorm:"column:city;type:varchar(255)" json:"city"`                 // 所在城市
	District     string    `gorm:"column:district;type:varchar(255)" json:"district"`         // 所在地区
}

// SysUserRole 用户与角色对应关系
type SysUserRole struct {
	ID     int64 `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	UserID int64 `gorm:"column:user_id;type:bigint(20)" json:"user_id"` // 用户ID
	RoleID int64 `gorm:"column:role_id;type:bigint(20)" json:"role_id"` // 角色ID
}

// User [...]
type User struct {
	ID             int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                // 主键
	Username       string    `gorm:"unique;column:username;type:varchar(50);not null" json:"username"`       // 登录名
	Password       string    `gorm:"column:password;type:varchar(100);not null" json:"password"`             // 登录密码
	NickName       string    `gorm:"column:nick_name;type:varchar(50)" json:"nick_name"`                     // 昵称
	UserPhoto      string    `gorm:"column:user_photo;type:varchar(100)" json:"user_photo"`                  // 用户头像
	UserSex        bool      `gorm:"column:user_sex;type:tinyint(1)" json:"user_sex"`                        // 用户性别，0：男，1：女
	AccountBalance int64     `gorm:"column:account_balance;type:bigint(20);not null" json:"account_balance"` // 账户余额
	Status         bool      `gorm:"column:status;type:tinyint(1);not null" json:"status"`                   // 用户状态，0：正常
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`           // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`           // 更新时间
}

// UserBookshelf 用户书架表
type UserBookshelf struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                  // 主键
	UserID       int64     `gorm:"unique_index:key_uq_userid_bookid;column:user_id;type:bigint(20);not null" json:"user_id"` // 用户ID
	BookID       int64     `gorm:"unique_index:key_uq_userid_bookid;column:book_id;type:bigint(20);not null" json:"book_id"` // 小说ID
	PreContentID int64     `gorm:"column:pre_content_id;type:bigint(20)" json:"pre_content_id"`                              // 上一次阅读的章节内容表ID
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

// UserBuyRecord 用户消费记录表
type UserBuyRecord struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                   // 主键
	UserID        int64     `gorm:"unique_index:key_userId_indexId;column:user_id;type:bigint(20);not null" json:"user_id"`    // 用户ID
	BookID        int64     `gorm:"column:book_id;type:bigint(20)" json:"book_id"`                                             // 购买的小说ID
	BookName      string    `gorm:"column:book_name;type:varchar(50)" json:"book_name"`                                        // 购买的小说名
	BookIndexID   int64     `gorm:"unique_index:key_userId_indexId;column:book_index_id;type:bigint(20)" json:"book_index_id"` // 购买的章节ID
	BookIndexName string    `gorm:"column:book_index_name;type:varchar(100)" json:"book_index_name"`                           // 购买的章节名
	BuyAmount     int       `gorm:"column:buy_amount;type:int(11)" json:"buy_amount"`                                          // 购买使用的屋币数量
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`                                       // 购买时间
}

// UserFeedback [...]
type UserFeedback struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"` // 主键id
	UserID     int64     `gorm:"column:user_id;type:bigint(20)" json:"user_id"`           // 用户id
	Content    string    `gorm:"column:content;type:varchar(512)" json:"content"`         // 反馈内容
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`     // 反馈时间
}

// UserReadHistory 用户阅读记录表
type UserReadHistory struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`                                  // 主键
	UserID       int64     `gorm:"unique_index:key_uq_userid_bookid;column:user_id;type:bigint(20);not null" json:"user_id"` // 用户ID
	BookID       int64     `gorm:"unique_index:key_uq_userid_bookid;column:book_id;type:bigint(20);not null" json:"book_id"` // 小说ID
	PreContentID int64     `gorm:"column:pre_content_id;type:bigint(20)" json:"pre_content_id"`                              // 上一次阅读的章节内容表ID
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}
