package main

import "fmt"

const (
    lcl = 122  // lowercase limit
    lcb = 97   // kowercase begins
    jump = 13  // jump
    upl = 90
    upb = 65
)


// Version 1 and i didnt like the code, really bad :(
// I will improve it
func rot13(message string) string {

    ordinals := []byte(message)
    for i, v := range ordinals {
        if v >= lcb && v <= lcl {
            if v + jump > lcl {
                ordinals[i] = (v + jump) - lcl + lcb - 1
            } else {
                ordinals[i] = v + jump
            }
        }

        if v >= upb && v <= upl {
            if v + jump > upl {
                ordinals[i] = (v + jump) - upl + upb - 1
            } else {
                ordinals[i] = v + jump
            }
        }

    }

    return string(ordinals)
}

func main() {
    message := "HELLO"
    fmt.Println(message)
    message = rot13(message)
    fmt.Println(message)

    message = "The Quick Brown Fox Jumps Over The Lazy Dog"
    fmt.Println(message)
    message = rot13(message)
    fmt.Println(message)



}
