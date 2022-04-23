package main

import (
	"github.com/tokizuoh/go-errors-verification/p1"
	"github.com/tokizuoh/go-errors-verification/p2"
)

func main() {
	print(p1.ErrInternal == p2.ErrInternal)
}
