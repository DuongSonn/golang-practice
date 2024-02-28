# Golang Practice

## Types

### Numbers

- Integer:

  - `uint8 ,uint16, uint64, int8, int16, unt31, unt64` (8,16,32,64 are how many bits each type uses)
  - `uint` (unsigned integer) contains positive numbers or zero
  - `byte` the same as `uint8`
  - `rune` the same as `int32`

- Floating-point number:

  - `float32, float64`
  - `complex64, complex128`: represent complex number with imaginary numbers
  - Larger size floating-point numbers increase its precision.
  - Can represent: NaN and positive, negative infinity

- String:

  - A space is also considered a character
  - String index start from 0.
  - Character is presented by byte => `fmt.Println("Hello"[1])` will print 101 (byte of e).
    Explain: get the character index 1 in string `Hello`

- Boolean:

## Variables

- `var x string = "Hello World"` or `x := "Hello World"`
- Variable name should start with a letter or \_. Go compiler doesn't care about name of a variable
- Scope: Variable exists within the nearest {} or block, including any nested curly braces but not outside of them
- Constants: `const x string = "Hello world"`
- Defining Multiple Variables:

```golang
var (
    a = 5
    b = 6
    c = 7
)
```

## Control Structures

- The `for` Statement:

```golang
func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }
}

func main() {
    for i + = 1; i <= 10; i++ {
        fmt.Println(i)
    }
}

func main() {
    for i, value := range x {

    }
}
```

- The `if` Statement:
- The `switch` Statement:

## Arrays, Slices, and Maps

### Arrays

- Is a numbered sequence of elements of a single type with a fixed length

```golang
    var x [5]int
    x := [5]int{1,2,3,4,5}
```

### Slices

- Is a segment of a array. But its length is allowed to change

```golang
    var x []float64
    // This creates a slice that is associated with an underlying float64 array of length 5
    x := make([]float64, 5)

    // This creates a slice with length of 5 that is associated with an underlying float64 array of length 10
    x := make([]float64, 5, 10)

    // This create a slice from index 0 -> 4 (5-1) from arr
    arr := [5]float64{1,2,3,4,5}
    x := arr[0:5]
    // Create slice form index 0 to end
    x := arr[0:]
    x := arr[0:len(arr)]
    x := arr[:]
    x := arr[:5]
```

- `append`: add elements onto the end of a slice. If there is not enough sufficient capacity => create new slice then add the new elements
- `copy`: copy all src to dst. If 2 slices have different length => smaller one will be used

### Maps

- Is an unordered collection of key-value pairs(dictionaries, hash tables)
- Map doesn't have fixed length

```golang
    x := make(map[string]int)
    x["1"] = 1

    x := map[string]string{
        "1": "1"
    }
```

## Functions

- Parameters names can be different
- Variables must be passed to functions
- Functions form a call stacks: Each time a function is called, we push it onto the call stack. Each time we return a function. we pop the last function off of the stack
- Return types can have names

```golang
    func f2() (r int) {
        return 1
    }
```

### Variadic Functions

```golang
    func add(args ...int) int {
        total := 0
        for _, v := range args {
            total += v
        }

        return total
    }

    func main() {
        fmt.Println(add(1,2,3))
        x := []int{1,2,3}
        fmt.Println(add(x...))
    }
```

- In above example, add allow to be called with multiple integers. This is called variadic parameter
- The `...` indicate it takes 0 or more int.
- Can also pass a slice of int to the function using `...`

### Closure

```golang

```

-
