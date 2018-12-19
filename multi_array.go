package main

import "fmt"


func main() {
    var a [3][3]int
    a[0][0] = 1
    fmt.Print(a)
    b := [3][3]int {
       {1, 33, 44},
       {3, 4, 55},
       {1, 44, 44},
    }
    fmt.Print(b) 
}
