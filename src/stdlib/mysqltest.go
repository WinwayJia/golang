package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gosexy/redis"
)

type Person struct {
	Id   int    "id"
	Name string "name"
	Sex  int    "sex"
}

var LOG *log.Logger

func main() {
	logfile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY, 0666)
	defer logfile.Close()
	LOG = log.New(logfile, "prefix", 0)
	
	fmt.Println(sql.Drivers())
//	db, err := sql.Open("mysql", "jiawh:111111@tcp(192.168.20.128:3306)/test?charset=utf8")
	db, err := sql.Open("mysql", "jiawh:111111@tcp(192.168.20.135:3307)/test?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	rows, err := db.Query("select * from user_1")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	defer rows.Close()
	r := redis.New()
	r.Connect("192.168.20.128", 6379)
	defer r.Close()

	for rows.Next() {
		var p Person
		rows.Scan(&p.Id, &p.Name, &p.Sex)
		str := fmt.Sprintf("%d:%s:%d", p.Id, p.Name, p.Sex)
		data, err := json.Marshal(p)
		if err != nil {
			continue
		}
		r.Set(str, string(data))
	}
	data, err := r.Get("12:2:2")
	fmt.Println(data)
	var pp Person
	err = json.Unmarshal([]byte(data), &pp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	fmt.Printf("%d:%s:%d\n", pp.Id, pp.Name, pp.Sex)
	LOG.Printf("error")
}
