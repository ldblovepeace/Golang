package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/GO-SQL-Driver/MySQL"
)

var Db *sql.DB
var sessionID int

func init() {
	var err error
	Db, err = sql.Open("mysql", "ldb:853126656@tcp(localhost:3306)/ldbsql?charset=utf8&parseTime=true ")
	//parseTime=true 将sql timestamp转换为time.Time格式

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Db.Stats())
	return
}

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func UserByEmail(email string) (user User, err error) {
	result, err := Db.Query("select id, uuid, name, email, password, created_at from users where email = $1", email)
	if err != nil {
		return
	}

	if err = result.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return
	}

	result.Close()
	return

}

func (session *Session) Check() (result bool, err error) {
	err = nil
	return true, err
}

func (user *User) CreateSession() (session Session) {
	sessionID++
	fmt.Println("sessionID now is:", sessionID)

	session.Id = sessionID
	session.UserId = user.Id
	session.CreatedAt = time.Now()
	session.Email = user.Email
	session.Uuid = user.Uuid

	stmt, _ := Db.Prepare("insert into sessions set id=?,uuid= ?,email= ?, user_id=?,created_at=?")
	stmt.Exec(sessionID, user.Uuid, user.Email, user.Id, time.Now())

	return session
}
