package main

import (
	"database/sql"
	"fmt"

	_ "github.com/GO-SQL-Driver/MySQL"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "ldb:853126656@tcp(localhost:3306)/ldbsql?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("init error")
	}
}

func (post *Post) Create() (err error) {
	stmt, err := db.Prepare("insert into post (content, author) values (?, ?)")
	if err != nil {
		fmt.Println("create stmt error")
		return
	}
	defer stmt.Close()
	stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Delete() (err error) {
	stmt, err := db.Prepare("delete from post where id = ?")
	if err != nil {
		fmt.Println("delete stmt error")
		return err
	}
	stmt.Exec(post.Id)
	defer stmt.Close()
	return
}

func (post *Post) Update() (err error) {
	_, err = db.Exec("update post set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	if err != nil {
		fmt.Println("update error")
	}
	return
}

func getPost(id int) (post Post, err error) {
	row := db.QueryRow("select * from post where id = ?", id)
	if err != nil {
		fmt.Println("post error")
		return
	}

	post = Post{}

	row.Scan(&post.Id, &post.Content, &post.Author)

	// fmt.Println("row:", row)
	return
}

func getPostsbyAuthor(author string) (posts []Post, err error) {
	rows, err := db.Query("select * from post where author = ?", author)
	if err != nil {
		fmt.Println("query error")
		return
	}

	p := Post{}

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Content, &p.Author)
		if err != nil {
			fmt.Println("scan error")
			return
		}
		posts = append(posts, p)
	}
	return
}

func main() {
	// Db, err := sql.Open("mysql", "ldb:853126656@tcp(localhost:3306)/ldbsql?charset=utf8&parseTime=true&loc=Local")
	// rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads order by created_at desc")

	// var Id, UserId int
	// var Uuid, Topic string
	// var CreatedAt time.Time

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for rows.Next() {
	// 		rows.Scan(&Id, &Uuid, &Topic, &UserId, &CreatedAt)
	// 		//这里要用& 不然结果写不到参数内，都为空
	// 		fmt.Println("id:", Id, "uuid:", Uuid, "topic:", Topic, "userid:", UserId, "created date:", CreatedAt)
	// 	}
	// }

	// post := Post{
	// 	Content: "Hello World",
	// 	Author:  "qwerty",
	// }

	// err := post.Create()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// t := time.Now().UnixNano()
	// post2, _ := getPost(2)

	// fmt.Println(post2.Id, post2.Content, post2.Author)
	// t = time.Now().UnixNano() - t
	// fmt.Println("time: ", t)

	posts, _ := getPostsbyAuthor("qwerty")

	for _, p := range posts {
		fmt.Println(p.Id, p.Content, p.Author)
	}
	// fmt.Println(posts)
}
