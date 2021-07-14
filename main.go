package main

import (
	"fmt"

	"github.com/lihakim/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	// definition, err := dictionary.Search("se")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	word := "hello"
	definition := "greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	newWordDef, err := dictionary.Search(word)
	fmt.Println(newWordDef)
	// 에러를 확인하기 위해 의도적으로 같은 단어 추가
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
