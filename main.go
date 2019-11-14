package main

import (
	_ "github.com/xiekch/cloudgo/service"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets/")))
	err := http.ListenAndServe(":8080", nil) // port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
