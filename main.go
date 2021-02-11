package main

import (
	"fmt"
	"net/http"
)

func main() {
	var counter int
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		counter += 1
		text := fmt.Sprintf("counter: %d", counter)
		w.Write([]byte(text))
		return
	})

	fmt.Println("server start at port 7000")
	http.ListenAndServe(":7000", nil)
}
