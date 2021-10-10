package follow

import (
	"fmt"
	"testing"
)

// このテストの中でmain()関数は実行できない、やりたいなら、main.goの結合テストの中で行う
func TestHello8(t *testing.T) {
	expect := "aaa"
	actual := "bbb"
	if expect != actual {
		t.Errorf("expect: %v, actual: %v", expect, actual)
	}
	fmt.Println("hello")
}
