package main

import (
    "fmt"
)

/* Main.
 */
func main() {
/*
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
*/

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

    /* Solution 2
     */
/*
    var counter int = 1
    type Records struct {
        account  []string
        uniqueid  int
    }
    var emailtable = make(map[string]*Records)
    var outputs2 = make(map[int]*Records)
    for _, account := range accounts {
        emails := account[1:]
        record := &Records{account: account, uniqueid: counter}
        
        for _, email := range emails {
            // There is a record, append emails to it
            if value, exists := emailtable[email]; exists == true {
                // Update the temporary record.
                record = value
                record.account = value.account
                record.uniqueid = value.uniqueid
            }

            // Append email if not already in list
            if InList(email, record.account) == false {
                record.account = append(record.account, email)
            }
            
            emailtable[email] = record
        }

        // 
        if _, found := outputs2[record.uniqueid]; found == false {
            outputs2[record.uniqueid] = record
        }

        counter += 1
    }

//    fmt.Println("Emailtable: ")
//    for k, v := range emailtable {
//        fmt.Printf("email=%s, value=%v\n", k, v)
//    }
    fmt.Println("Outputs2:")
    for _, v := range outputs2 {
        fmt.Printf("%v\n", v)
    }
    // End Solution 2 //
*/
}
