package main

import (
	"database/sql"
	"fmt"

	_ "github.com/GO-SQL-Driver/MySQL"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "ldb:853126656@tcp(localhost:3306)/ldbsql?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("insert userinfo set username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("ldb", "IT", "2018-10-16")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("ldbupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
