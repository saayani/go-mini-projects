package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build MapHandler using the mux as the fallback

	pathsToUrls := map[string]string{
		"/panic":  "https://medium.com/swlh/simple-guide-to-panic-handling-and-recovery-in-golang-72d6181ae3e8",
		"/struct": "https://yourbasic.org/golang/structs-explained/",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on: 8080")

	http.ListenAndServe(":8080", mapHandler)
}

// func defaultMux() {
// }
