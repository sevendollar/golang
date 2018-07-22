package main

import (
    "fmt"
)

type my_error struct {
    err string
}

func (e *my_error) Error() string {
    return fmt.Sprintf("%s", e.err)
}

func err_test(condition float64) (result string, err error) {
    switch {
    case condition >= 1:
        result = "good result!"
        return
    case condition <= 0:
        result = "bad result!"
        err = &my_error{"error: value can't be low then equal to 0"}
    }
    return
}

func main() {
    v, err := err_test(-115)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(v)
    }
}
