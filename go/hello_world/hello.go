package main

import "fmt"

import "rsc.io/quote"

func main() {
    var i int
    i = 42
    j := 42.0
    var k float32 = 42.0
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    fmt.Printf("i: value %v, type %T\n", i, i)
    fmt.Printf("j: value %v, type %T\n", j, j)
    fmt.Printf("i: value %v, type %T\n", k, k)
}
