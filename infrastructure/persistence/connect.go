package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var  db *sql.DB

func CreateAcess() (int64, error) {
  _ , err := sql.Open("mysql", "root:@tcp(192.168.1.7:3306)/poker")
  if err != nil {
    panic(err)
  }
  return 1,err
}

func InsertIntoDB(query string)(int64,error){
  _ , err := db.Exec(query)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  return 1, nil
}

func main() {
  CreateAcess()
  fmt.Println(InsertIntoDB("USE poker"))
}
