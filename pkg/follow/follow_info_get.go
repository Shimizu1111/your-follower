package follow

import (
	"fmt"
	"net/http"
)

type PostForm struct {
}

// 対象アカウントのフォロー情報を取得する
func FollowInfoGet(w http.ResponseWriter, req *http.Request) {
	// curlをたたく、フォローの情報を取得するAPI
	// curlの結果、スライス、配列形式
	follow := []string{"aaa", "bbb"}
	fmt.Println(follow)

	// このフォームの内容をテストしたいなら、インタフェース化して外だしする
	// このままではテストできない、テスト時にリクエストに値を入れるようにする(フォームに値が入ると同じ)
	PostForm := &bbb{
		Name: req.FormValue("name"),
	}
	fmt.Println("postForm =", PostForm)
	ResultOutput(w, req)
	fmt.Println("Follow")
	fmt.Fprint(w, "helloafdfd")
}
