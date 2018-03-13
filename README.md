Packages

+ package main defines the stand alone executable
+ What main does is what the program does

Imports

+ must only define the imports you need
+ imports must follow the package declaration
+ followed by declarations of functions, variables, constants and types (func, var, const, type)

Functions

+ delcaration func
+ name of the function
+ parameter list
+ result list
+ body of the function


+ Does not require semicolons at the end of statements or declarations

+ newlines folwing certain declarations converted to semicolons
    + the { has to be on the same line as the function declaration
    + x + y ; newline is permitted after but not before the +

GOFMT

+ go tools gofmt applies formatting to go code
+ go get golang.org/x/tools/cmd/goimport

Command Line Arguments

+ The os package provides functions and other values for dealing with the operating system in a platform-independent fashion
+ os.Args is a slice of strings
+ slices will be more in depth later but quickly, dynamically sized sequence s of array elements 
    + accessed by s[i]
    + contiguous subsequence s[m:n]
    + number of elements len(s)
    + os.Args[1:len(os.Args)]
    + os.Args[1:]

+ while loop
+ infinite loop

1. Modify the echo program to also print os.Args[0]. What is it?
2. Modify the echo program to print the index and value of each of the arguments, one per line
3. How can we rewrite the above programs using the Strings std library ?
4. Experiment to measure the difference in running time between our potentially inefficent versions and the one using the strings package 

resources: 

https://golang.github.io/dep/
https://play.golang.com/
https://godoc.org/