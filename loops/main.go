package main

import "fmt"

func main() {
    fmt.Println("----- Basic For Loop -----")
	// for i := 0; i < 5; i++ {
    for i := range 5 {
        fmt.Println("i =", i)
    }

    fmt.Println("\n----- Using break -----")
    for i := 1; i <= 5; i++ {
        if i == 3 {
            fmt.Println("Breaking the loop at i =", i)
            break
        }
        fmt.Println("i =", i)
    }

    fmt.Println("\n----- Using continue -----")
    for i := 1; i <= 5; i++ {
        if i == 3 {
            fmt.Println("Skipping i =", i)
            continue
        }
        fmt.Println("i =", i)
    }

    fmt.Println("\n----- Infinite Loop with break -----")
    counter := 0
    for {
        if counter == 3 {
            fmt.Println("Breaking infinite loop at counter =", counter)
            break
        }
        fmt.Println("counter =", counter)
        counter++
    }

    fmt.Println("\n----- Using goto (not recommended) -----")
    i := 0
start:
    if i < 3 {
        fmt.Println("goto i =", i)
        i++
        goto start
    }
}
