package user

import (
	"fmt"
	//	"crypto/md5"
	"encoding/json"
	"errors"
	"msg"

	"github.com/gosexy/redis"
)

type UserInfo struct {
	ID       uint64
	Name     string
	NickName string
	Password string
	Age      int
	Addr     string
}

type UserChat struct {
	Info UserInfo
	In   chan msg.MsgInfo
}

func (self *UserInfo) Register() error {
	r := redis.New()
	if err := r.Connect("192.168.20.128", 6379); err != nil {
		return err
	}
	defer r.Close()
	//	md := md5.New()
	//	md.Sum(self.Password)
	//	md.Write()
	id := fmt.Sprintf("%v", self.ID)

	value, _ := json.Marshal(self)
	r.HSet("user:info", id, value)
	return nil
}

func (self *UserInfo) Login() (bool, error) {
	r := redis.New()
	if err := r.Connect("192.168.20.128", 6379); err != nil {
		return false, err
	}
	defer r.Close()
	var user UserInfo
	id := fmt.Sprintf("%v", self.ID)
	value, err := r.HGet("user:info", id)
	if err != nil {
		fmt.Println("HGet Error:", id, value)
		return false, err
	}
	fmt.Println("HGet", id, value)
	json.Unmarshal([]byte(value), &user)
	if self.Password != user.Password {
		fmt.Println(self.Password, user.Password)
		return false, errors.New("incorrect password")
	}

	return true, nil
}
