package matcher

import "github.com/yshujie/suggest/internal/suggest-core/model"

type DictionaryMatcher struct {
}

// Match 根据关键词匹配词典
func (m *DictionaryMatcher) Match(keyword string) (model.Dictionary, error) {
	return nil, nil
}
