package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/atotto/clipboard"
	cc "golang.design/x/clipboard"
)

func TestCtrl(t *testing.T) {
	clipboard.WriteAll("中文")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)
	// 不支持 中文
}

func TestChina(t *testing.T) {
	for range time.NewTicker(time.Second).C {
		text, err := clipboard.ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(text, "----")
	}
	// 不支持 中文

}

func TestCCCC(t *testing.T) {
	ch := cc.Watch(context.TODO(), cc.FmtText)
	for data := range ch {
		// print out clipboard data whenever it is changed
		println(string(data))
	}

}
