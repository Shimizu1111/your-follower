package follow

import (
	"fmt"
	"html/template"
	"net/http"
)

type bbb struct {
	Name string
}

func ResultOutput(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("../../template/result.html")
	if err != nil {
		panic(err.Error())
	}

	values := &bbb{
		Name: req.FormValue("name"),
	}

	fmt.Println(values)

	if err := t.Execute(w, values); err != nil {
		panic(err.Error())
	}
}
