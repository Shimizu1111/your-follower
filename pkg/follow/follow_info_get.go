package follow

import (
	"fmt"
	"net/http"
)

// 対象アカウントのフォロー情報を取得する
func FollowInfoGet(w http.ResponseWriter, req *http.Request) {

	values := &bbb{
		Name: req.FormValue("name"),
	}
	fmt.Println("values2", values)
	ResultOutput(w, req)
	fmt.Println("Follow")
}
