package follow

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// このテストの中でmain()関数は実行できない、やりたいなら、main.goの結合テストの中で行う
// 対象アカウントのフォロー情報を取得するためのテスト
// この機能とmain()を用いたテストは、結合テストになるためmain_test.goファイルで行う

func TestHello8(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/follow", nil)
	// w http.ResponseWriter, req *http.Request
	FollowInfoGet(response, request)
	actual := response.Body.String()
	expect := "bbb"
	if expect != actual {
		t.Errorf("expect: %v, actual: %v", expect, actual)
	}
	fmt.Println("hello")
	fmt.Println("response = ", response)
	fmt.Println("request = ", request)
}
