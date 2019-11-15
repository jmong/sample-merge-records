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

func InList(needle string, haystack []string) bool {
    for _, v := range haystack {
        if v == needle {
            return true
        }
    }
    return false
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
    fmt.Printf("Accounts: %v\n", accounts)

    // Solution 1 //
    var outputs [][]string
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
    // End Solution 1 //

    // Solution 2 - work-in-progress //
    type Records struct {
        account []string
    }
    var emailtable = make(map[string]*Records)
    var outputs2 []*Records
    for _, account := range accounts {
        name, emails := account[0], account[1:]
        for _, email := range emails {
            // There is a record, append emails to it
            if val, ok := emailtable[email]; ok == true {
                acct := val.account
                if InList(email, acct) == false {
                    acct = append(acct, email)
                }
                val.account = acct
            } else {  // No existing record, create one
                acct := []string{name}
                acct = append(acct, emails...)
                rec := &Records{account: acct}
                emailtable[email] = rec
                outputs2 = append(outputs2, rec)
            }
        }
    }

    fmt.Println("Emailtable: ")
    //for k, v := range emailtable {
    //    fmt.Printf("email=%s, val=%v\n", k, v)
    //}
    fmt.Println("Outputs2:")
    for _, v := range outputs2 {
        fmt.Printf("%v\n", v)
    }
    // End Solution 2 //
}
