package main

import (
	"fmt"
	"net/http"
	"os"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	hostName, _ := os.Hostname()
	data := fmt.Sprintf("Hello %s\n", hostName)
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
