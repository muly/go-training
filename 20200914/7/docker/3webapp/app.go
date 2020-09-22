package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!!!!!!!!!!!"))
}

// docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
//		docker build -t webapp .
// docker run to create the container from the specified image (using tag in below example)
// 		docker run --rm -p 8000:8080 webapp
