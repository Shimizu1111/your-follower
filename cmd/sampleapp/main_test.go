package main

import (
	"fmt"
	"testing"
)

func TestHello6(t *testing.T) {
	expect := "aaa"
	actual := "bbb"
	if actual != expect {
		t.Errorf("actual: %v, expected: %v", actual, expect)
	}
	fmt.Println("helloTestdaaa")
	Hello5()
	//main()
}
