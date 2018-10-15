package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ldblovepeace/test/web/action"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "HOMEPAGE") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", homepage)          //设置访问的路由
	http.HandleFunc("/login", action.Login) //设置访问的路由
	http.HandleFunc("/upload", action.Upload)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
