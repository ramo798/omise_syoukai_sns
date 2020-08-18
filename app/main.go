package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World2222")
}

func main() {
	fmt.Print("Server Start")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
