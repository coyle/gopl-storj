# Program Structure
    Basic structural elements of a Go program

## Names

    A name begins with a letter or an underscore
    May have any number of aditional letters, digitds and underscores

    * case matters (heapSort vs. HeapSort)

    Go has 25 keywords

    ~ three dozen predeclared names (int, true, new)
        * Constants, Types, Functions 
        * not reserved
    
    If an entity is declared within a function it is local to that function

    If declared outside the function it is visible to all files of the package it belongs to

        * Uppercase declarations are exported
        * package names are always lowercase

    No limit on name length but idiomatic to use descriptive names for larger scopes
        * (i vs. theLoopIndex)

    Camel Case is preferred 
    acronyms should be capitalized (escapeHtml vs escapeHTML)

## Declarations

    declaration names a program entity and specifies some or all of its properties.

    Four major kinds of declarations:
        1. var
        2. const
        3. type
        4. func
    
    A Go program is stored in one or more files that end in .go

    A Go file begins with a package declaration, followed by import declarations, then a sequence of package-level declarations of types, variables, constants, and functions.

    https://play.golang.com/p/lEsrLUSWhdO

    https://play.golang.com/p/Qkbxt-mf-da

## Variables

    A var declaration creates a variable of a particular type, attaches a aname to it and sets its initial value.
     * var **name type** = **expression**
     * either the **type** or **expression** may be omitted but not both
     * If **type** is omitted then the type is determined from the expression
     * If the **expression** is omitted then the initial value is the zero value for the declared type
     * 0, false, "", nil, []int{0} 

     * no such thing as an uninitialized variable
     ```Go
        var s string
        fmt.Println(s) // ""
    ```

    ```Go
        var i, j, k int // int, int, int
        var b, f, s = true, 2.3, "four" // bool, float64, string
    ```

    * package level variables are initialized before main begins
    * local variables are initialized as their declarations are encountered during function execution

    * set of variables can be initialized by a function that returns variables
    ```Go
        var f, err = os.Open(**name**)
    ```

### Short Variable Declarations
    * **name** := **expression**

    * type of **name** is determined by the type of **expression**

    ```Go
        anim := gif.GIF{loopCount: nframes}
        freq := rand.Float64() * 3.0
        t := 0.0
    ```

    * Used to declare the majority of local variables
    * var typically is reserved for declarations that need an explicit type that differs from that of the initializer expression
    
    ```Go
        i := 0                      \\ an int
        var boiling float64 = 100   \\ a float64
    ```

    multiple variables maybe declared
    ```Go
        i, j := 0, 1
    ```
    but should only be used when they help readability
    * := is a declaration
    * = is an assignment
    ```Go
        i, j = j, i // swap values of i and j
    ```

    may be used for calls to functions that return two or more variables
    ```Go
        f, err := os.Open(**name**)
    ```

    * One subtle  but important point is that a short declaration does not necessarily **declare** all the variables on its left hand side. If some were already declared in the same lexical block then it acts as an assignment

    ```Go
        in, err := os.Open(**name**) // declares both in and err
        out, err := os.Create(**name**) // declares out but assigns a value to the existing err
    ```

    * a sort declaration must declare at least one new variable
    ```Go
        f, err := os.Open(**name**)
        f, err := os.Create(**name**) // won't compile compile error: no new variables
    ```

### Pointers

    A pointer value is the **address** of a variable.
        * The location at which a value is stored
        * not every value has an address but every variable does
        * can read or update the value of a variable indirectly without even knowing the name of the variable

    If a variable is declared:
    ```Go
        var x int
    ```

    the expression **&x**  ("the address of **x**") yields a pointer to an integer variable that is a value of type *int ("pointer to int")
    * if it is called "p" we say that "p points to x" or "p contains the address to x"
    * the variable to which p points is written as *p. 
    * the expression *p yields the value of that variable, an int
    * since *p denotes a variable it may be placed on the left hand side of an assignment, in which case the assignment updates the variable

    ```Go
        x := 1
        p := &x         // p, of type *int, points to x
        fmt.Println(*p) // "1"
        *p = 2          // equivalent to x = 2
        fmt.Println(x)  // "2"
    ```

    * Each component of an aggregate type (field on a struct or element of an array) is a variable and therefore has an address

    * zero value of a pointer is nil
    * p != nil is true if **p** points to a variable
    * pointers are comparable, they are equal only if they both point to the same address/variable, or both are nil

    ```Go
        var x, y int
        fmt.Println(&x == &x, &x == &y, &x == nil) 
    ```

    * if a function returns the address to a local variable, the local variable will stay in existence even after the function has returned

    ```Go
        var p = f()

        func f() *int {
            v := 1
            return &v
        }
    ```
    * each call to f returns a distinct value

    ```Go
        fmt.Println(f() == f())  // "false"
    ```

    * Since a pointer contains the address of a variable, passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed.

    ```Go
        func incr(p *int) int {
            *p++ // increments what p points to; does not change p
            return *p
        }

        v := 1
        incr(&v)                // side effect: v is now 2
        fmt.Println((incr(&v))  // "3" and v is now 3
    ```
    

### The new Function

Another way to create a variable is with the function **new**

* The function new(T) creates an unnamed variable of type T, initializes it to the zero value of T and returns its address which is a value of type *T

* no different than an ordinary local variable whose address is taken except that there's no need to invent (declare) a dummy name and we can use new(T) in an expression

* only syntactic convenience
```Go
    func newInt() *int {
        return new(int)
    }

    func newInt() *int {
        var dummy int
        return &dummy
    }
```

* Each call to new returns a distinct variable with a unique address

```Go
    p := new(int)
    q := new(int)
    fmt.Println( p == q) // "false"
```

* since new is a predeclared function, not a keyword, it is possible to redefine the name for something else

```Go
    func delta(old, new int) int {return new - old}
```

### Lifetime of Variables

* The lifetime of a variable is the interval of time it exists as the program executes

* package level variable = entire time the progrma executes
* local variables = created each time the declaration statement is executed, lives until unreachable
* function parameters are local variables
* compiler determines which variables are allocated to heap or stack
* while GC is a tremendous help, does not relieve you of the burden of thinking about memory
* aware of lifetime of variables to allow GC to free memory
* unnecessary pointers to short lived objects with long lived objects, prevents GC from reclaiming short lived objects


## Packages and Files

* packages in Go serve the same purpose as libraries or modules in other languages
* supporting modularity, encapsulation, seperate compilation, and reuse
source code for a packe resides in one or more .go files
* files of github.com/coyle/gopl-storj are stored in $GOPATH/src/coyle/gopl-storj

* each package serves as a seperate namespace for its declarations

* packages allow for information hiding

### Imports

### Package Imitialization

## Scope
