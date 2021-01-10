package dao

import "gorm.io/gorm"

type BaseDao struct {
	DB *gorm.DB
}

// 这么做是为了兼容事务操作
func (bd *BaseDao) GetDb() *gorm.DB {
	// 根据db指针是否为空来选择一个原始DB或是一个事务Session等.
	if bd.DB == nil {
		return GetDb()
	}
	return bd.DB
}
