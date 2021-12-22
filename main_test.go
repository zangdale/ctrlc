package main

import (
	"fmt"
	"testing"

	"github.com/atotto/clipboard"
)

func TestCtrl(t *testing.T) {
	clipboard.WriteAll("中文")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)

	// Output:
	// 中文
}
