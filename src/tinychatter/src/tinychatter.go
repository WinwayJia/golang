package main

import (
	"fmt"
	"msg"
	"os"
	"time"
	"user"
)

type UserMap map[uint64]user.UserChat

var g_usermap UserMap

func recv_msg(u *user.UserChat) {
	var msg msg.MsgInfo
	fmt.Println("routine", u)
	for {
		msg = <-u.In
		fmt.Println(msg)
	}
}

func main() {
	userInfo := user.UserInfo{123,
		"NULL",
		"LOL",
		"Password",
		20,
		"Earth",
	}

	g_usermap = make(map[uint64]user.UserChat)

	err := userInfo.Register()
	if err != nil {
		os.Exit(1)
	}
	userChat := user.UserChat{userInfo, make(chan msg.MsgInfo, 1024)}
	ret, err := userInfo.Login()
	//	u := user.UserInfo{}
	if ret {
		u, ok := g_usermap[userInfo.ID]
		if !ok {
			g_usermap[userInfo.ID] = userChat
			u = userChat
		}
		fmt.Println("Login success")
		go recv_msg(&u)
	} else {
		fmt.Println("Login failed")
	}
	for {
		time.Sleep(time.Second * 3)
		msg := msg.MsgInfo{123, 123, "Hello"}
		//		fmt.Println(userChat)
		userChat.In <- msg
		//		fmt.Println("new loop")
	}
}
