package main

import "fmt"


func Staircase (levels int) {

    n := levels
    for n > 0 {
        for i := 0; i < levels; i++ {
            if i < n - 1 {
                fmt.Print(" ")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Println()
        n = n - 1
    }
}


func main() {
    var levels int
    fmt.Scanf("%d", &levels)
    Staircase(levels)
}
