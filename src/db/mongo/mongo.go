package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	url = "127.0.0.1"
)

// Person struct
type Person struct {
	Name  string
	Phone string
}

func insert(c *mgo.Collection) {
	err := c.Insert(map[string]interface{}{"id": "99090"})
	if err != nil {
		painc(err)
	}

	user := struct {
		*Person
		ID       int64  `json:"id"`
		Nickname string `json:"nickname"`
	}{
		Person: &Person{
			Name:  "name",
			Phone: "186++++",
		},
		ID:       90000,
		Nickname: "hello",
	}
	user.Phone = "+86000000"

	err := c.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

}

func query(c *mgo.Collection) {
	result := struct {
		*Person
		ID       int64  `json:"id"`
		Nickname string `json:"nickname"`
	}{}
	err := c.Find(bson.M{"nickname": "hello"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Name)
}

func main() {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("mgo.Dial failed.")
	}
	err = session.Ping()
	if err != nil {
		log.Print(err)
	}

	database := session.DB("test")
	collection := database.C("user")

	insert(collection)
	query(collection)

	defer session.Close()
}
