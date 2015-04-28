package main

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	codisHost := "192.168.188.190:19000"
	c, err := redis.DialTimeout("tcp", codisHost, 5*time.Second, 5*time.Second, 5*time.Second)
	checkErr(err)
	defer c.Close()

	key := "foo"
	value, err := redis.Bytes(c.Do("GET", key))
	checkErr(err)
	fmt.Println(string(value))
}
