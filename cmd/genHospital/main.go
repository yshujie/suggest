package main

import (
	"github.com/yshujie/suggest/internal/suggest-core/manager"
	"github.com/yshujie/suggest/internal/suggest-core/model"
)

func init() {
	code := model.WordBankCode(model.Hospital)

	if _, ok := interface{}(code).(model.WordBankCode); !ok {
		panic("code is not a word bank code")
	}

	wordBackManager := manager.NewWordBackManager(code)

	wordBack, err := wordBackManager.CreateWordBack()
	if err != nil {
		panic("create word bank failed")
	}

	err = wordBackManager.InitWordBank(wordBack)
	if err != nil {
		panic("init word bank failed")
	}
}

func main() {

}
