package chain

/*
	假设我们现在有个校园论坛，由于社区规章制度、广告、法律法规的原因需要对用户的发言进行
敏感词过滤。如果被判定为敏感词，那么这篇铁则将会被封禁。
*/

// SensitiveWordFilter 敏感词过滤器，判断是否是敏感词
type SensitiveWordFilter interface {
	Filter(content string) bool
}

type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// 添加过滤器
func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// 执行过滤
func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		//如果发现敏感词直接返回结果
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// 广告
type AdSensitiveWordFilter struct{}

// Filter 实现过滤算法
func (f *AdSensitiveWordFilter) Filter(content string) bool {
	return false
}

// 政治敏感词汇
type PoliticalWordFilter struct{}

// Filter 实现过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	return true
}
