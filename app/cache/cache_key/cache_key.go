package cache_key

const (
	IndexClickBankBookKey = "indexClickBankBookKey"
	IndexNewBookKey       = "indexNewBookKey"
	IndexUpdateBookKey    = "indexUpdateBookKey"
	IndexBookKeyTime  	  = 300000 // 5分钟刷新一次, 单位ms
)
