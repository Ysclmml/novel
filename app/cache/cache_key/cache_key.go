package cache_key

const (
	IndexClickBankBookKey = "indexClickBankBookKey"
	IndexNewBookKey       = "indexNewBookKey"
	IndexUpdateBookKey    = "indexUpdateBookKey"
	IndexBookKeyTime      = 300000                 // 5分钟刷新一次, 单位ms
	IndexBookSettingsKey  = "indexBookSettingsKey" // 首页小说设置
)
