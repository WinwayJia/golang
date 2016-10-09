package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Person struct {
	Name string
	Age  int
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	indexFile, err := os.Open("html/index.html")
	if err != nil {
		fmt.Println(err)
	}
	index, err := ioutil.ReadAll(indexFile)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {

		}
		fmt.Println("Client subscribed")
		myPerson := Person{
			Name: "websocket",
			Age:  10,
		}
		for {
			if myPerson.Age < 40 {
				data, err := json.Marshal(myPerson)
				if err != nil {
					break
				}
				err = conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					break
				}
			} else {
				break
			}
		}
		conn.Close()
		fmt.Println("Client unsubscirbed")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})
	http.ListenAndServe(":3000", nil)
}
