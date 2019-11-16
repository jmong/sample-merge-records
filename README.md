# sample-merge-records

# Usage

Run it.<br>
(your output will differ slightly, eg- ordering of the records
and emails might be different)
```
$ go run -v .
<omitted>
Inputs: [[John johnA@mail.com johnB@mail.com johnC@mail.com] [John johnD@mail.com johnE@mail.com] [John johnA@mail.com johnF@mail.com] [Mary maryA@mail.com maryB@mail.com] [John johnE@mail.com johnG@mail.com johnH@mail.com] [Mary maryC@mail.com maryD@mail.com] [Mary maryC@mail.com maryE@mail.com]]

// Start Solution 1 //
Outputs:
[John johnA@mail.com johnF@mail.com johnB@mail.com johnC@mail.com]
[John johnE@mail.com johnG@mail.com johnH@mail.com johnD@mail.com]
[Mary maryA@mail.com maryB@mail.com]
[Mary maryC@mail.com maryE@mail.com maryD@mail.com]
// End Solution 1 //

// Start Solution 2 //
Outputs:
[John johnA@mail.com johnB@mail.com johnC@mail.com johnF@mail.com]
[John johnD@mail.com johnE@mail.com johnG@mail.com johnH@mail.com]
[Mary maryA@mail.com maryB@mail.com]
[Mary maryC@mail.com maryD@mail.com maryE@mail.com]
// End Solution 2 //
```

# Efficiency

TL;DR
> Solution 2 is much more efficient than Solution 1.

Solution1 iterates thru each input record ("current record"), 
compares all of the email addresses in the current record to every email 
addresses in all previous output records. 
`Merge()` is used to compare 2 lists of email addresses to find intersections and merge them.
It is inefficient because it compares every email address in one list with every email 
address in the other list, that is, at worse it can be O(m*n).

On the other hand, Solution2 is more efficient because it uses hashing to keep tracking
what email addresses has been seen, as a way to determine if there is an intersection.
Hash lookup is really efficient, that is, O(1).

See a sample benchmark of those 2 solutions side-by-side.<br>
Solution2 is almost 2x faster than Solution1.
```
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkRunSolution1-12    	  500000	      3035 ns/op	    2304 B/op	      29 allocs/op
BenchmarkRunSolution2-12    	 1000000	      1796 ns/op	    1207 B/op	      13 allocs/op
PASS
ok  	<omitted>  	3.377s
```

# Author

Joe Mong
