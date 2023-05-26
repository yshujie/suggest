package model

// WordBank 词库，是一类数据的全集，包含一个词库代码和一组词典
type WordBank struct {
	Code       string
	Dictionary []Dictionary
}
