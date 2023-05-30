package manager

import "github.com/yshujie/suggest/internal/suggest-core/model"

// WordBankManager 词库管理器
type WordBankManager struct {
	code        string           `json:"code,omitempty"`
	dictionarys model.Dictionary `json:"dictionarys,omitempty"`
}

// CreateWordBack 创建词库
func (m *WordBankManager) CreateWordBack(code string) (model.WordBank, error) {
	return model.WordBank{}, nil
}

// InitWordBank 初始化词库
func (m *WordBankManager) InitWordBank(wordBank model.WordBank) error {
	return nil
}

// DestroyWordBack 销毁词库
func (m *WordBankManager) DestroyWordBack(bank model.WordBank) error {
	return nil
}
