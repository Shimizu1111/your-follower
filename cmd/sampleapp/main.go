package main

import (
	"fmt"
	"html/template"
	"net/http"

	"example.com/integration"
	"example.com/sub"
	"example.com/sub/follow"
)

type aaa struct {
	Name string
	Age  int
}

func main() {
	test := "aaaa"
	fmt.Println("hello")
	sub.Sub()
	fmt.Println(test)
	http.HandleFunc("/", mainHandler)
	// 暫定的に今ここに置いているが、/follow、のPOSTを受け取った関数の中でこのhandleをするのでおそらく変える
	http.HandleFunc("/result", follow.ResultOutput)
	// この処理は、フォームで入力された内容をPOSTで受け取る処理、画面出力はしない
	http.HandleFunc("/follow", follow.FollowInfoGet)
	defer http.ListenAndServe(":8000", nil)
	follow.FollowerInfoCompare()
	follow.FollowerInfoGet()
	// テストファイルの動作確認用,integration
	integration.Hello()
	integration.Hello()
	Hello5()
	// フィールドに値をセット
	p2 := &aaa{
		Name: "typeですよ",
		Age:  23,
	}
	fmt.Println(p2)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../template/index.html")
	if err != nil {
		panic(err.Error())
	}
	p2 := &aaa{
		Name: "typeですよ2",
		Age:  25,
	}
	if err := t.Execute(w, p2); err != nil {
		panic(err.Error())
	}
}

func Hello5() {
	fmt.Println("helloTestFunc")
}
