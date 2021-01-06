package model

import "novel/app/utils/time"

// BookComment 小说评论表
type BookComment struct {
	ID             int64         `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id"`                                      // 主键
	BookId         int64         `gorm:"unique_index:key_uq_bookid_userid;column:book_id;type:bigint(20)" json:"book_id"`               // 小说ID
	CommentContent string        `gorm:"column:comment_content;type:varchar(512)" json:"comment_content"`                               // 评价内容
	ReplyCount     int           `gorm:"column:reply_count;type:int(11)" json:"reply_count"`                                            // 回复数量
	AuditStatus    bool          `gorm:"column:audit_status;type:tinyint(1)" json:"audit_status"`                                       // 审核状态，0：待审核，1：审核通过，2：审核不通过
	CreateTime     time.JsonTime `gorm:"column:create_time;type:datetime" json:"create_time"`                                           // 评价时间
	CreateUserID   int64         `gorm:"unique_index:key_uq_bookid_userid;column:create_user_id;type:bigint(20)" json:"create_user_id"` // 评价人
}

func (book *BookComment) TableName() string {
	return "book_comment"
}

// BookCommentReply 小说评论回复表
type BookCommentReply struct {
	ID           int64         `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`     // 主键
	CommentID    int64         `gorm:"column:comment_id;type:bigint(20)" json:"comment_id"`         // 评论ID
	ReplyContent string        `gorm:"column:reply_content;type:varchar(512)" json:"reply_content"` // 回复内容
	AuditStatus  bool          `gorm:"column:audit_status;type:tinyint(1)" json:"audit_status"`     // 审核状态，0：待审核，1：审核通过，2：审核不通过
	CreateTime   time.JsonTime `gorm:"column:create_time;type:datetime" json:"create_time"`         // 回复用户ID
	CreateUserID int64         `gorm:"column:create_user_id;type:bigint(20)" json:"create_user_id"` // 回复时间
}

func (book *BookCommentReply) TableName() string {
	return "book_comment_reply"
}