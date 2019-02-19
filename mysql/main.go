package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

const (
	DB_Driver = "root:fushihao@tcp(127.0.0.1:3306)/skate?charset=utf8"
)

func main() {
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("insert new_table set idnew_table=?")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("123")
	if err != nil {
		panic(err)
	}
}
