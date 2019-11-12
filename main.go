package main

import (
    "fmt"
)

// Merge compares two arrays and determine if there are any elements that are 
// the same (ie- intersect) and creates an union of the two arrays.
// A union is all non-repeating elements. 
// For example, union of {"a", "b", "c"} and {"c", "d", "e"} is {"a", "b", "c", "d", "e"}.
// @credit https://siongui.github.io/2018/03/10/go-set-of-all-elements-in-two-arrays/
func Merge(a, b []string) (bool, []string) {
    var intersect bool = false
    m := make(map[string]bool)

    for _, item := range a {
        m[item] = true
    }

    for _, item := range b {
        if _, ok := m[item]; !ok {
            a = append(a, item)
        } else {
            intersect = true
        }
    }

    return intersect, a
}

func main() {
    // More records added here from the original problem set.
    // The following accounts should merge:
    // 1. First and third
    // 2. Second and fifth
    // 3. Sixth and seventh
    var accounts = [][]string{
        []string{"John", "johnA@mail.com", "johnB@mail.com", "johnC@mail.com"},
        []string{"John", "johnD@mail.com", "johnE@mail.com"},
        []string{"John", "johnA@mail.com", "johnF@mail.com"},
        []string{"Mary", "maryA@mail.com", "maryB@mail.com"},
        []string{"John", "johnE@mail.com", "johnG@mail.com", "johnH@mail.com"},
        []string{"Mary", "maryC@mail.com", "maryD@mail.com"},
        []string{"Mary", "maryC@mail.com", "maryE@mail.com"},
    }
    var outputs [][]string

    fmt.Printf("Accounts: %v\n", accounts)
    for _, account := range accounts {
        var intersect bool = false
        var union []string
        for oidx, output := range outputs {
            // Intersection found, replace current outputs record with the merged emails
            if intersect, union = Merge(account[1:], output[1:]); intersect == true {
                newrecord := []string{account[0]}
                newrecord = append(newrecord, union...)
                outputs[oidx] = newrecord
                break
            }
        }

        // No intersection found, add as a new record in outputs
        if intersect == false {
            outputs = append(outputs, account)
        }
    }

    fmt.Printf("Outputs: %v\n", outputs)
}
