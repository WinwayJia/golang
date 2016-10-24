package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
	"encoding/json"
)

const (
	hash_key = "film_info"
	zset_key = "rating"
)

func queryRows(db *sql.DB) {
	rows, err := db.Query("select title, film_id, rental_rate from film;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	buffer := bytes.NewBuffer(make([]byte, 128))
	columns, _ := rows.Columns()
	for _, v := range columns {
		buffer.WriteString(v + " ")
	}
	log.Println(buffer.String())

	var name string
	var id int
	var rate float64
	for rows.Next() {
		rows.Scan(&name, &id, &rate)
		log.Printf("name: %s id: %d, rate: %.2f\n", name, id, rate)
		writeToRedis(name, id, rate)
	}

}

func insertRows(db *sql.DB) {
	tx, _ := db.Begin()
	//rows, err := db.Query(fmt.Sprintf("insert into user_info(user_name, user_id) values ('%s', %d);", "xml", 3))
	rows, err := db.Query("insert into user_info(user_name, user_id) values ($1, $2);",
		"xml", 3)
	if err != nil {
		log.Println(err, rows)
	}

	tableName := "user_info"
	columnName := "user_name"
	columnID := "user_id"
	db.Exec(fmt.Sprintf("insert into %s(%s, %s) values($1, $2)", tableName, columnName, columnID),
		"protobuff", 10)
	tx.Commit()
}

func writeToRedis(name string, id int, rate float64) error {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
		return err;
	}
	redisConn := redis.NewConn(conn, 30 * time.Second, 30 * time.Second)
	s := struct {
			Name string `json:"name"`
			ID int `json:"id"`
		} {
			Name: name,
			ID: id,
		}
	b, _ := json.Marshal(s)
	_, err = redisConn.Do("HSET", hash_key, id, string(b))
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = redisConn.Do("ZADD", zset_key, rate, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:111111@127.0.0.1:5432/dvdrental")
	//db, err := sql.Open("postgres", "user=postgres dbname=test sslmode=verify-full")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	//insertRows(db)

	queryRows(db)
}
