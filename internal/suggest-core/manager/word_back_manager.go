package manager

import (
	"github.com/yshujie/suggest/internal/suggest-core/model"
	"os"
	"path/filepath"
)

// WordBankManager 词库管理器
type WordBankManager struct {
	code string `json:"code,omitempty"`
}

// NewWordBackManager 创建词库管理器
func NewWordBackManager(code model.WordBankCode) *WordBankManager {
	return &WordBankManager{
		code: string(code),
	}
}

// CreateWordBack 创建词库
func (m *WordBankManager) CreateWordBack() (model.WordBank, error) {
	return model.WordBank{
		Code: m.code,
	}, nil
}

// InitWordBank 初始化词库
func (m *WordBankManager) InitWordBank(wordBank model.WordBank) error {
	file, err := m.openSourceDataFile()
	if err != nil {
		return err
	}

	// 创建词典管理器
	dictionaryManager := NewDictionaryManager(wordBank)

	// 字典 code 列表
	dictionaryCodes := []model.DictionaryCode{model.HashDictionaryCode, model.TernaryTreeCode}

	// 循环 dictionaryCodes，创建词典
	for _, code := range dictionaryCodes {
		dictionary, err := dictionaryManager.CreateDictionary(code)
		if err != nil {
			return err
		}

		err = dictionaryManager.InitDictionary(dictionary, file)
		if err != nil {
			return err
		}

		wordBank.AddDictionary(dictionary)
	}

	return nil
}

// openSourceDataFile 打开词库文件
func (m *WordBankManager) openSourceDataFile() (*os.File, error) {
	// 打开 code 对应的词库文件
	file, err := os.Open(m.buildFilePath())
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return file, nil
}

// buildFilePath 构建词库文件路径
func (m *WordBankManager) buildFilePath() string {
	return filepath.Join("data", m.code, m.code+".txt")
}

// DestroyWordBack 销毁词库
func (m *WordBankManager) DestroyWordBack(bank model.WordBank) error {
	return nil
}
