// 4.3.3 exercise: Given a list of names, you need to organize each name within a slice based on its length. http://www.golangbootcamp.com/book/collection_types
package main

import "fmt"

func main() {
    var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
        "Emma", "Isabella", "Emily", "Madison",
        "Ava", "Olivia", "Sophia", "Abigail",
        "Elizabeth", "Chloe", "Samantha",
        "Addison", "Natalie", "Mia", "Alexis"}

    // iterates over the variable "names" and pick the maximum length.
    max_len := func(x []string) (result int) {
        for _, v := range names {
            if len(v) > result {
                result = len(v)
            }
        }
        return
    }(names)
    //  initialize the slice of slice of string.
    result := make([][]string, max_len)
    fmt.Println("before:", result)

    result = func(x *[][]string, y []string) [][]string {
        xx := *x
        for _, v := range y {
            xx[len(v)-1] = append(xx[len(v)-1], v)
        }
        *x = xx
        return *x
    }(&result, names)
    fmt.Println("after:", result)
}
