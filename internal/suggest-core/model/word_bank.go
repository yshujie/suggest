package model

// WordBankCode 定义 WordBack 的 code 类型，其中有 Hospital, Child, Patient, Doctor
type WordBankCode string

const (
	Hospital WordBankCode = "hospital"
	Child    WordBankCode = "child"
	Patient  WordBankCode = "patient"
	Doctor   WordBankCode = "doctor"
)

// WordBank 词库，是一类数据的全集，包含一个词库代码和一组词典
type WordBank struct {
	Code         string
	Dictionaries []Dictionary
}

// AddDictionary 添加词典
func (b *WordBank) AddDictionary(dictionary Dictionary) {
	b.Dictionaries = append(b.Dictionaries, dictionary)
}
