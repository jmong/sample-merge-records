package main

/* Merge compares two arrays and determine if there are any elements that are 
 * the same (ie- intersect) and creates an union of the two arrays.
 * A union is all non-repeating elements. 
 * For example, union of {"a", "b", "c"} and {"c", "d", "e"} is {"a", "b", "c", "d", "e"}.
 * @credit https://siongui.github.io/2018/03/10/go-set-of-all-elements-in-two-arrays/
 */ 
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

/* RunSolution1
 */
func RunSolution1(inputs [][]string) [][]string {
    var outputs [][]string
    for _, record := range inputs {
        var intersect bool = false
        var union []string
        for oidx, output := range outputs {
            // Intersection found, replace current outputs record with the merged emails
            if intersect, union = Merge(record[1:], output[1:]); intersect == true {
                newrecord := []string{record[0]}
                newrecord = append(newrecord, union...)
                outputs[oidx] = newrecord
                break
            }
        }

        // No intersection found, add as a new record in outputs
        if intersect == false {
            outputs = append(outputs, record)
        }
    }

    return outputs
}
