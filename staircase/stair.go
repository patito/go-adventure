package main

import "fmt"


func Staircase (num int) {

    n := num
    for n > 0 {
        for i := 0; i < num; i++ {
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

    x := "we"  * 2
    fmt.Scanf(x)
    var num int
    fmt.Scanf("%d", &num)
    Staircase(num)
}
