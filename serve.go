package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
