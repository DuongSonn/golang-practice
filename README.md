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
    func makeEvenGenerator() func() uint {
        i := uint(0)
        return func() (ret uint) {
            ret = i
            i += 2
            return
        }
    }

    func main() {
        x := 0
        increment := func() int {
            x++
            return x
        }
        fmt.Println(increment()) // Print 1
        fmt.Println(increment()) // Print 2

        nextEven := makeEvenGenerator()
        fmt.Println(nextEven()) // 0
        fmt.Println(nextEven()) // 2
        fmt.Println(nextEven()) // 4
    }
```

- It is possible to create functions inside of functions

### Recursion

### defer

- defer is often used when resources need to be freed in some way
- if function has multiple return statements, defer func will happen before any of them
- defer functions are run even if a runtime panic occurs
- Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

### panic and recover

```golang
    func main() {
        defer func() {
            str := recover()
            fmt.Println(str)
        }()

        panic("PANIC")
    }
```

- `panic` indicates a programmer error or exceptional condition that there is no way to recover from it
- we use panic function to create a runtime error
- recover stops the panic and returns the value that was passed to the call panic

### Pointers

```golang
    func zero(xPtr *int) {
        *xPtr = 0 // Store the int 0 in the memory location xPtr refers to.
    }
    func main() {
        x := 5
        zero(&x)
        fmt.Println(x) // x is 0

        y := new(int)
        zero(y)
        fmt.Println(*y)
    }
```

- pointers reference a location in memory where a value is stored rather than the value itself
- The `*` give access to the value the pointer points to
- The `&` find the address of a variable
- The `new` return a pointer to the type it takes as an argument

## Structs and interfaces

### Structs

```golang
    type  Circle struct {
        x,y,r float64
    }

    var c Circle
    c := new(Circle) // The new function return a pointer to the struct (*Circle)
    c := &Circle{0,0,0}

    // area is a method of Circle
    func (c *Circle) area() float64 {
        return math.Pi * c.r*c.r
    }

    type Person struct {
        Name string
    }


    type Android struct {
        Person Person // This is called named field
        Model string
    }
    a := new(Android)
    a.Person.Talk()

    type Android struct {
        Person // This is anonymous field
        Model string
    }
    a := new(Android)
    a.Talk()
```

- Pointer are often used with structs so functions can modify their contents
- `x,y,z` are called fields
- `method` are associated with types => its behavior is specific to a type
- `method` have an implicit receiver => its allow methods to access and modify the properties of the object
- `struct` defines fields
- `interface` defines a `method` set

## Packages

## Testing

## Concurrency

- It helps handle one or more tasks simultaneously

### Goroutines

- Is a function capable of running concurrently with other functions.

```golang
    func f(n int) {
        for i := 0; i < 10; i++ {
            fmt.Println(n, ":", i)
        }
    }

    func main() {
        go f(0)
        var input string
        fmt.Scanln(&input)
    }
```

### Mutex

- Provides a concurrent-safe way to express exclusive access to these shared resources. It will create critacl sessions for the resourcé
- Critical sections are so named because they reflect a bottleneck in your program. It is somewhat expensive to enter and exit a critical section, and so generally people attempt to minimize the time spent in critical sections
=> Solutions we use `sync.RWMutex`
- In `sync,RWMutext` you can request a lock for reading, in which case you will be granted access unless the lock is being held for writing. This means that an arbitrary number of readers can hold a reader lock so long as nothing else is holding a writer lock.

### Cond

- In some cases, you want goroutine to stop at a certain condition then contiunue executing affter a signal is received => You use `Cond`
```
c := sync.NewCond(&sync.Mutex{})
queue := make([]interface{}, 0, 10)
removeFromQueue := func(delay time.Duration) {
  time.Sleep(delay)
  c.L.Lock()
  queue = queue[1:]
  fmt.Println("Removed from queue")
  c.L.Unlock()
  c.Signal()
}
for i := 0; i < 10; i++{
  c.L.Lock()
  for len(queue) == 2 {
    c.Wait()
  }
  fmt.Println("Adding to queue")
  queue = append(queue, struct{}{})
  go removeFromQueue(1*time.Second)
  c.L.Unlock()
}
```
- `Wait` doesn’t just block, it suspends the current goroutine, allowing other goroutines to run on the OS thread. Upon entering Wait, `Unlock` is called on the `Cond` variable’s `Locker`, and upon exiting `Wait`, `Lock` is called on the `Cond` variable’s `Locker`
- Internally, the run‐time maintains a FIFO list of goroutines waiting to be signaled; `Signal` finds the goroutine that’s been waiting the longest and notifies that
- `Broadcast` sends a signal to all goroutines that are waiting
- `Once`  is a type that utilizes some sync primitives internally to ensure that only one call to `Do` ever calls the function passed in—even on different goroutines. You can use it to guard against multiple initialization.
- `Pool`

### Channels

```golang
    func pinger(c chan string) {
        for i := 0; ; i++ {
            c <- "ping"
        }
    }
    func ponger(c chan string) {
        for i := 0; ; i++ {
            c <- "pong"
        }
    }

    func printer(c chan string) {
        for {
            msg := <- c
            fmt.Println(msg)
            time.Sleep(time.Second * 1)
        }
    }

    func main() {
        var c chan string = make(chan string)

        // This will take turn print ping and pong
        go pinger(c)
        go ponger(c)
        go printer(c)

        var input string
        fmt.Scanln(&input)
    }
```

- Channels provide a way for 2 goroutines to communicate with each other and synchronize their execution
- Channel is represented with the keyword `chan` followed by the type passed to the channel
- `c <-` is to send data to the channel
- `<- c` is to receive data from the channel
- Channel is blocking. This means that any goroutine that attempts to write to a channel that is full will wait until the channel has been emptied, and any goroutine that attempts to read from a channel that is empty will wait until at least one item is placed on it.
- You can check the value read is generated by another process or from close channel like this: `salutation, ok := <-stringStream`
- Closing a channel is also one of the ways you can signal multiple goroutines simultaneously. If you have n goroutines waiting on a single channel, instead of writing n times to the channel to unblock each goroutine, you can simply close the channel.

```golang
    func pinger(chan <- ) // pinger is only allowed to send to c
    func printer(<- chan) // printer is only allowed to receive from c
```

- We can restrict channel to either send of receive data

```golang
    func main() {
        c1 := make(chan string)
        c2 := make(chan string)
        go func() {
            for {
                c1 <- "from 1"
                time.Sleep(time.Second * 2)
            }
        }()
        go func() {
            for {
                c2 <- "from 2"
                time.Sleep(time.Second * 3)
            }
        }()
        go func() {
            for {
                select {
                    case msg1 := <- c1:
                        fmt.Println(msg1)
                    case msg2 := <- c2:
                        fmt.Println(msg2)
                    case <- time.After(time.Second):
                        fmt.Println("timeout")
                }
            }
        }()

        var input string
        fmt.Scanln(&input)
    }
```

- `select` work like `switch` for channel
- `select` pick the 1st channel that is ready and receive from it (or sends to it). If more than 1 channels are ready, it randomly pick 1. If none is ready, it blocks until one becomes available
- `time.After` after duration (1 second) of waiting, the function will print timeout. this prevent from waiting forever
- `default` case will happen if none of the channel is ready
- Channels are synchronous; both side of the channel will wait until the other side is ready

```golang
    c := make(chan int, 1)
```

- Unbufferd Channel is synchronous: sending and receiving a message will perform one after another.
- Buffered Channel is asynchronous; sending or receiving a message will not wait unless the channel is full. If the channel is full, the sending will wait until there is room for more
- `b := make(chan int, 0)` is ab unbuffered channel. An unbuffered channel has a capacity of zero and so it’s already full before any writes
