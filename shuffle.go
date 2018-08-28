package shuffle

import (
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
