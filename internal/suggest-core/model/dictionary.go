package model

// Dictionary 词典，是一类数据的不同实现，比如：HashDictionary、TernaryTreeDictionary
type Dictionary interface {
	GetCode() string
	GetSize() int
	GetLen() int
	GetLastUpdateTime() int64

	query(keyword string) ([]Entry, error)
	addEntry(entry Entry) error
	updateEntry(entry Entry) error
}

// HashDictionary 哈希词典，是一类数据的实现，包含一个词典代码、大小、长度、词条列表和最后更新时间
type HashDictionary struct {
	Code           string
	Size           int
	Len            int
	Entry          []Entry
	LastUpdateTime int64
}

func (d *HashDictionary) GetCode() string {
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

func (d *HashDictionary) query(keyword string) ([]Entry, error) {
	return nil, nil
}

func (d *HashDictionary) addEntry(entry Entry) error {
	return nil
}

func (d *HashDictionary) updateEntry(entry Entry) error {
	return nil
}

// TernaryTreeDictionary 三叉树词典，是一类数据的实现，包含一个词典代码、大小、长度、词条列表和最后更新时间
type TernaryTreeDictionary struct {
	Code           string
	Size           int
	Len            int
	Entry          []Entry
	LastUpdateTime int64
}

func (d *TernaryTreeDictionary) GetCode() string {
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

func (d *TernaryTreeDictionary) query(keyword string) ([]Entry, error) {
	return nil, nil
}

func (d *TernaryTreeDictionary) addEntry(entry Entry) error {
	return nil
}

func (d *TernaryTreeDictionary) updateEntry(entry Entry) error {
	return nil
}
