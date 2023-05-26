package manager

import "github.com/yshujie/suggest/internal/suggest-core/model"

type DictionaryEditor struct {
	dic model.Dictionary
}

// AddEntry 添加词条
func (m *DictionaryEditor) AddEntry(entry model.Entry) error {
	return nil
}

// RemoveEntry 移除词条
func (m *DictionaryEditor) RemoveEntry(entry model.Entry) error {
	return nil
}

// UpdateEntry 更新词条
func (m *DictionaryEditor) UpdateEntry(entry model.Entry) error {
	return nil
}
