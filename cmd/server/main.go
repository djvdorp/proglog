package main

import (
	"log"

	"github.com/djvdorp/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":9080")
	log.Fatal(srv.ListenAndServe())
}
