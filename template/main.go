package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type Item struct {
// 	Title     string    `json:"title"`
// 	CreatedAt time.Time `json:"created_at"`
// }

type address struct {
	Message interface{} `json:"message"`
	Results []bbb       `json:"results"`
	Status  int         `json:"status"`
}

type bbb struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Kana1    string `json:"kana1"`
	Kana2    string `json:"kana2"`
	Kana3    string `json:"kana3"`
	Prefcode string `json:"prefcode"`
	Zipcode  string `json:"zipcode"`
}

func main() {
	resp, err := http.Get("https://zipcloud.ibsnet.co.jp/api/search?zipcode=1350062")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("jsonの中身は", string(body))

	var data2 address

	if err := json.Unmarshal(body, &data2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("data2= ", data2)
	for _, item := range data2.Results {
		fmt.Printf("%s %s\n", item.Address1, item.Address2)
	}
}

// func main2() {
// 	resp, err := http.Get("http://qiita.com/api/v2/users/snaka/items?page=1&per_page=10")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp)
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var data []Item // nil slice
// 	// data := make([]Item, 0) のように要素数0の slice としても良い
// 	fmt.Println("bodyのないようは", string(body))

// 	if err := json.Unmarshal(body, &data); err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, item := range data {
// 		fmt.Printf("%s %s\n", item.CreatedAt, item.Title)
// 	}
// }
