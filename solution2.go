package main

/* 
 */
type RecordObj struct {
    // 
    record    []string

    // Id to uniquely identify this particular instance. 
    uniqueid  int
}

/* InList checks if needle is in the list of haystack.
 */
func InList(needle string, haystack []string) bool {
    for _, v := range haystack {
        if v == needle {
            return true
        }
    }
    return false
}

/*
 */
func RunSolution2(inputs [][]string) map[int]*RecordObj {
    var (
        // Used to generate unique ids
        uniqueid int  = 1
        
        // Hash where the key is an email address and value holds the pointer 
        // to a RecordObj. We know that an email address is seen again in another
        // record if the value holds a RecordObj pointer.
        emailhash    = make(map[string]*RecordObj)
        
        //
        outputs      = make(map[int]*RecordObj)
    )

    for _, record := range inputs {
        emails := record[1:]
        temprecord := &RecordObj{record: record, uniqueid: uniqueid}
        
        for _, email := range emails {
            // There is an existing record
            if value, exists := emailhash[email]; exists == true {
                // Replace with the existing record
                // and update its values
                temprecord = value
                temprecord.record = value.record
                temprecord.uniqueid = value.uniqueid
            }

            // Append email if not already in list
            if InList(email, temprecord.record) == false {
                temprecord.record = append(temprecord.record, email)
            }
            
            emailhash[email] = temprecord
        }

        // 
        if _, found := outputs[temprecord.uniqueid]; found == false {
            outputs[temprecord.uniqueid] = temprecord
        }

        uniqueid += 1
    }

    return outputs
}
