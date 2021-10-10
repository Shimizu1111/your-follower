package main

import (
	"fmt"
	"testing"
)

func TestHello6(t *testing.T) {
	// expectの値をフォームからの入力を想定する
	// expect := "@shimizu"    これがフォーム側から入力されるようにする
	expect := ""
	actual := "@shimizu"
	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
	fmt.Println("helloTestdaaa")
	Hello5()
	//main()
}
