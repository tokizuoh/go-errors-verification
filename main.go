package main

import (
	"fmt"

	"github.com/tokizuoh/go-errors-verification/p1"
	"github.com/tokizuoh/go-errors-verification/p2"
)

// func New(text string) error {
// 	return errorString(text)
// }

// type errorString string

// func (e errorString) Error() string {
// 	return string(e)
// }

func main() {

	// errors で作成した同じ文字列のエラー比較
	fmt.Println(p1.ErrInternal == p2.ErrInternal) // false

	// 自作エラー (errorStringを構造体からstringに変更)で作成した同じ文字列のエラー比較
	// fmt.Println(New("Internal error") == New("Internal error"))
}
