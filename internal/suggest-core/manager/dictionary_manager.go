package manager

import "github.com/yshujie/suggest/internal/suggest-core/model"

// DictionaryManager 词典管理器
type DictionaryManager struct {
	wBank model.WordBank
}

func NewDictionaryManager(wBack model.WordBank) *DictionaryManager {
	return &DictionaryManager{
		wBank: wBack,
	}
}

// CreateDictionary 创建词典
func (m *DictionaryManager) CreateDictionary(code string) (model.Dictionary, error) {
	return nil, nil
}

// InitDictionary 初始化词典
func (m *DictionaryManager) InitDictionary(dic model.Dictionary) error {
	return nil
}

// DestroyDictionary 销毁词典
func (m *DictionaryManager) DestroyDictionary(dic model.Dictionary) error {
	return nil
}
