package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir //strings.Replace(dir, "\\", "/", -1)
}

//test template function
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Println(getCurrentDirectory())
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New(`../HTML/testfunction.html`).Funcs(funcMap)
	//路径如果不在同一目录下 浏览器会有错误 具体什么原因？？？
	//问题就处在这个路径上，路径不能有‘/’
	//纯文本文件的字符编码未声明。如果该文件包含 US-ASCII 范围之外的字符，
	//该文件将在某些浏览器配置中呈现为乱码。该文件的字符编码需要在传输协议层声明，
	//或者在文件中加入一个 BOM（字节顺序标记）。
	t, _ = t.ParseFiles(`../HTML/testfunction.html`)
	t.Execute(w, time.Now())
}

//for test
func test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../HTML/testfunction.html")
	t.Execute(w, r)
}

func main() {
	http.HandleFunc("/", homepage)          //设置访问的路由
	http.HandleFunc("/login", action.Login) //设置访问的路由
	http.HandleFunc("/logins", logins)      //test session
	http.HandleFunc("/upload", action.Upload)
	http.HandleFunc("/setMessage", setMassage)
	http.HandleFunc("/getMessage", getMessage)
	http.HandleFunc("/process", process)
	http.HandleFunc("/test", test)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
