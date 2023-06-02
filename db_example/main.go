package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    config := mysql.Config {
        User: os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "gotest",
    }
    var err error
    db, err = sql.Open("mysql", config.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

}
