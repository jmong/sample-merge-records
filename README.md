# sample-merge-records

# Usage

Build it.
```
go build -v
```

Run it (actual output included here).
```
./sample-merge-records
Accounts: [[John johnA@mail.com johnB@mail.com johnC@mail.com] [John johnD@mail.com johnE@mail.com] [John johnA@mail.com johnF@mail.com] [Mary maryA@mail.com maryB@mail.com] [John johnE@mail.com johnG@mail.com johnH@mail.com] [Mary maryC@mail.com maryD@mail.com] [Mary maryC@mail.com maryE@mail.com]]
Outputs: [[John johnA@mail.com johnF@mail.com johnB@mail.com johnC@mail.com] [John johnE@mail.com johnG@mail.com johnH@mail.com johnD@mail.com] [Mary maryA@mail.com maryB@mail.com] [Mary maryC@mail.com maryE@mail.com maryD@mail.com]]
```

# Author

Joe Mong
