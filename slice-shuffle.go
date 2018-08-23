package main

import (
    "fmt"
    "math/rand"
    "time"
)

func shuffle(x []int) {
    rand.Seed(time.Now().UnixNano())
    for i := range x {
        j := rand.Intn(i + 1)
        x[i], x[j] = x[j], x[i]
    }
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    for i := 0; i < 15; i++ {
        shuffle(a)
        fmt.Println(a)
    }
}
