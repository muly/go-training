// demonstrate a simple app with docker (with small image size)

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

//// multi-step container build:
// docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
//		docker build -t docker2 .
// docker run to create the container from the specified image (using tag in below example)
// 		docker run --rm docker2

// https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/ #Part 3: Compile!
// vs
// https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2 #Way 2: Build Outside your Dockerfile
