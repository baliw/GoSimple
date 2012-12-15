
package main

import (
	"fmt"
	"net/http"
	"time"
	"flag"
	)

func main() {
	var port *int = flag.Int("port", 8080, "Port to listen on. Default=8080")
	var path *string = flag.String("path", "./", "Path to serve files from.  Default=./")

	fileserver := http.FileServer(http.Dir(*path))

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", *port),
		Handler: fileserver,
		ReadTimeout: 30*time.Second,
		WriteTimeout: 30*time.Second,
	}
	fmt.Printf("Listening on port %d\n", *port)
	fmt.Printf("Serving from: %s\n", *path)
	err := srv.ListenAndServe()
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}


