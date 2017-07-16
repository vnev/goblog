package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/index"
	"blog/blog"
	"log"
)

func main() {
	router := httprouter.New()
	router.GET("/", index.HomepageHandler)
	router.GET("/blog", blog.IndexHandler)
	router.GET("/blog/:id", blog.PostHandler)

	log.Fatal(http.ListenAndServe(":3000", http.Handler(router)))
}
