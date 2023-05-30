package manager

import (
	"bufio"
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

	dictionaryManager := NewDictionaryManager(wordBank)

	hashDictionary, err := dictionaryManager.CreateDictionary("hash")
	ternaryTreeDictionary, err := dictionaryManager.CreateDictionary("ternaryTree")

	// 读取词库文件中的词条
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entryStr := scanner.Text()

		// 将词条添加到 Hash 词典中
		hashDictionary.AddEntry(model.Entry{
			Word: entryStr,
		})

		// 将词条添加到 TernaryTree 词典中
		ternaryTreeDictionary.AddEntry(model.Entry{
			Word: entryStr,
		})
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
