package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Static/")))
	//http.Handle("/", http.FileServer(http.Dir("/.Static")))

	fmt.Println("Http File Server Running...")

	log.Fatal(http.ListenAndServe(":8003", nil))
}

