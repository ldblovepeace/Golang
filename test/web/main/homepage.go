package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/ldblovepeace/test/web/action"
	"github.com/ldblovepeace/test/web/common/session"
	_ "github.com/ldblovepeace/test/web/common/session/providers/memory"
)

var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
}

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

func logins(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		cookie, err := r.Cookie("gosessionid")

		if err == nil {
			t, _ := template.ParseFiles("../HTML/afterlogin.html")
			t.Execute(w, sess.GetbySessionID(cookie.Value))
		} else {
			t, _ := template.ParseFiles("../HTML/logins.html")
			w.Header().Set("Content-Type", "text/html")
			t.Execute(w, sess.Get("username"))
		}
	} else {
		sess.Set("username", r.Form["username"])
		value, _ := sess.Get("username").(interface{})
		fmt.Println(value)
		http.Redirect(w, r, "/", 302)
	}
}

//set flash cookie
func setMassage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello world")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
	fmt.Fprintln(w, "set cookie success")
}

//get flash cookie
func getMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		fmt.Fprintln(w, "no message")
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc) //将name为flash的cookie消除了（maxage = -1）
		msg, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(msg))
	}
}

func main() {
	http.HandleFunc("/", homepage)          //设置访问的路由
	http.HandleFunc("/login", action.Login) //设置访问的路由
	http.HandleFunc("/logins", logins)      //test session
	http.HandleFunc("/upload", action.Upload)
	http.HandleFunc("/setMessage", setMassage)
	http.HandleFunc("/getMessage", getMessage)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
