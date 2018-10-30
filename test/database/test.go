package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/GO-SQL-Driver/MySQL"
)

func main() {
	Db, err := sql.Open("mysql", "ldb:853126656@tcp(localhost:3306)/ldbsql?charset=utf8&parseTime=true&loc=Local")
	rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads order by created_at desc")

	var Id, UserId int
	var Uuid, Topic string
	var CreatedAt time.Time

	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			rows.Scan(&Id, &Uuid, &Topic, &UserId, &CreatedAt)
			//这里要用& 不然结果写不到参数内，都为空
			fmt.Println("id:", Id, "uuid:", Uuid, "topic:", Topic, "userid:", UserId, "created date:", CreatedAt)
		}
	}
}
