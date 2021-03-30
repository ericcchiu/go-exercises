Problem Description: Two-fer
---
Two-fer or 2-fer is short for two for one. One for you and one for me.

Given a name, return a string with the message:
```Text
One for X, one for me.
```

Where X is the given name.
However, if the name is missing, return the string:

```Text
One for you, one for me.
Here are some examples:
```

`Name	String to return`
Alice	One for Alice, one for me.
Bob	One for Bob, one for me.
One for you, one for me.
Zaphod	One for Zaphod, one for me.

Coding the solution
---
Look for a stub file having the name two_fer.go and place your solution code in that file.

Running the tests
---
To run the tests run the command go test from within the exercise directory.

If the test suite contains benchmarks, you can run these with the --bench and --benchmem flags:
```Golang
go test -v --bench . --benchmem
```
Keep in mind that each reviewer will run benchmarks on a different machine, with different specs, so the results from these benchmark tests may vary.