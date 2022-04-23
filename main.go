package main

import (
	"errors"
	"fmt"

	"github.com/tokizuoh/go-errors-verification/p1"
	"github.com/tokizuoh/go-errors-verification/p2"
)

func main() {
	// errors で作成した同じ文字列のエラー比較
	fmt.Println(errors.New("Internal error") == errors.New("Internal error")) // false

	// 自作エラー (errorStringを構造体からstringに変更)で作成した同じ文字列のエラー比較
	fmt.Println(p1.ErrInternal == p2.ErrInternal) // true
}
