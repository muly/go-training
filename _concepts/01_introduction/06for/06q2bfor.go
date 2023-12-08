package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 10; j++ {
            if i*j > 10 {
                continue
            }
            fmt.Printf("%d x %d = %d\n", i, j, i*j)
        }
    }
}