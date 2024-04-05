package util

import (
	"fmt"
	"github.com/bytedance/sonic"
	"strconv"
)

type User struct {
	Name string
}

func Rs() string {
	u := User{Name: strconv.Itoa(1234)}
	brr, err := sonic.Marshal(u)
	if err != nil {
		fmt.Println("这里面的数据不正常")
	} else {
		fmt.Println("这里面的数据正常")
	}
	return string(brr)
}
