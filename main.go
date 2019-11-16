package main

import (
    "fmt"
)

/* Main.
 */
func main() {
    fmt.Printf("Inputs: %v\n\n", Inputs)

    // Solution 1
    fmt.Println("// Start Solution 1 //")
    fmt.Println("Outputs:")
    for _, entry := range RunSolution1(Inputs) {
        fmt.Printf("%v\n", entry)
    }
    fmt.Println("// End Solution 1 //")

    fmt.Println()

    // Solution 2
    fmt.Println("// Start Solution 2 //")
    fmt.Println("Outputs:")
    for _, entry := range RunSolution2(Inputs) {
        fmt.Printf("%v\n", entry.record)
    }
    fmt.Println("// End Solution 2 //")
}
