package model

// Dictionary 词典，是一类数据的不同实现，比如：HashDictionary、TernaryTreeDictionary
type Dictionary struct {
	Code           string
	Size           int
	Len            int
	Entry          []Entry
	lastUpdateTime int64
}
