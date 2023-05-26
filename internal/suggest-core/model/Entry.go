package model

// Entry 词条，词库中的最小元素，包含索引和对象
type Entry struct {
	Index  EntryIndex
	Object EntryObject
}

// EntryIndex 词条索引
type EntryIndex interface{}

// EntryObject 词条对象
type EntryObject interface {
	GetCode() string
	ToString() string
	ToJson() string
}
