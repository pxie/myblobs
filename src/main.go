/**
 * Created by pxie on 12/10/13.
 */
package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	io.WriteString(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	fmt.Printf("Hello world!")
}
