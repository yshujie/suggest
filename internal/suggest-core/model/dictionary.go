package model

type DictionaryCode string

const (
	HashDictionaryCode DictionaryCode = "hash"
	TernaryTreeCode    DictionaryCode = "ternary_tree"
)

// Dictionary 词典，是一类数据的不同实现，比如：HashDictionary、TernaryTreeDictionary
type Dictionary interface {
	GetCode() WordBankCode
	GetSize() int
	GetLen() int
	GetLastUpdateTime() int64

	Query(keyword string) ([]Entry, error)
	AddEntry(entry Entry) error
	UpdateEntry(entry Entry) error
}

// HashDictionary 哈希词典，是一类数据的实现，包含一个词典代码、大小、长度、词条列表和最后更新时间
type HashDictionary struct {
	Code           WordBankCode
	Size           int
	Len            int
	Entry          []Entry
	LastUpdateTime int64
}

// NewHashDictionary 创建一个哈希词典
func NewHashDictionary(code WordBankCode) *HashDictionary {
	return &HashDictionary{
		Code: code,
		Size: 0,
		Len:  0,
	}
}

func (d *HashDictionary) GetCode() WordBankCode {
	return d.Code
}

func (d *HashDictionary) GetSize() int {
	return d.Size
}

func (d *HashDictionary) GetLen() int {
	return d.Len
}

func (d *HashDictionary) GetLastUpdateTime() int64 {
	return d.LastUpdateTime
}

func (d *HashDictionary) Query(keyword string) ([]Entry, error) {
	return nil, nil
}

func (d *HashDictionary) AddEntry(entry Entry) error {
	return nil
}

func (d *HashDictionary) UpdateEntry(entry Entry) error {
	return nil
}

// TernaryTreeDictionary 三叉树词典，是一类数据的实现，包含一个词典代码、大小、长度、词条列表和最后更新时间
type TernaryTreeDictionary struct {
	Code           WordBankCode
	Size           int
	Len            int
	Entry          []Entry
	LastUpdateTime int64
}

// NewTernaryTreeDictionary 创建一个三叉树词典
func NewTernaryTreeDictionary(code WordBankCode) *TernaryTreeDictionary {
	return &TernaryTreeDictionary{
		Code: code,
		Size: 0,
		Len:  0,
	}
}

func (d *TernaryTreeDictionary) GetCode() WordBankCode {
	return d.Code
}

func (d *TernaryTreeDictionary) GetSize() int {
	return d.Size
}

func (d *TernaryTreeDictionary) GetLen() int {
	return d.Len
}

func (d *TernaryTreeDictionary) GetLastUpdateTime() int64 {
	return d.LastUpdateTime
}

func (d *TernaryTreeDictionary) Query(keyword string) ([]Entry, error) {
	return nil, nil
}

func (d *TernaryTreeDictionary) AddEntry(entry Entry) error {
	return nil
}

func (d *TernaryTreeDictionary) UpdateEntry(entry Entry) error {
	return nil
}
