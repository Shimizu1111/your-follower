package follow

import (
	"fmt"
	"testing"
)

// このテストの中でmain()関数は実行できない、やりたいなら、main.goの結合テストの中で行う
// 対象アカウントのフォロー情報を取得するためのテスト
// この機能とmain()を用いたテストは、結合テストになるため、main_test.goファイルで行う
func TestHello8(t *testing.T) {
	expect := "aaa"
	actual := "bbb"
	if expect != actual {
		t.Errorf("expect: %v, actual: %v", expect, actual)
	}
	fmt.Println("hello")
}
