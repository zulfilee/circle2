package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Spectro!")
}

func main() {
	http.HandleFunc("/hello-backend", handler)
	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
//     SECRET=(samplesecret)

// http://mainsite.sample.com
// Adding an extra comment
// Adding an extra comment
}
