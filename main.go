
package main

import (
	"fmt"
	"net/http"
	"time"
	)

func main() {
	fileserver := http.FileServer(http.Dir("/opt/htdocs"))

	srv := &http.Server{
		Handler: fileserver,
		ReadTimeout: 30*time.Second,
		WriteTimeout: 30*time.Second,
	}
	fmt.Printf("Listening on port 80...\n")
	err := srv.ListenAndServe()
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}


