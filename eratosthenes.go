package main

import "fmt"

func main() {
    var primes [10000]bool
    for i := 2; i < len(primes); i++ {
        if primes[i] { continue }
        for j := 2 * i; j < len(primes); j += i{
            primes[j] = true 
        }
        fmt.Println(i)
    }
}
