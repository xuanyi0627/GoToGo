package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello xx!") //这个写入到w的是输出到客户端的
}

func main()  {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}