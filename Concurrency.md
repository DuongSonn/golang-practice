# Concurrency In Golang

## Introductions

### Race Conditions

- A race condition occurs when 2 or more operations must execute in the correct order but the program has not been written so that this order is guaranteed to be maintained
- Data race is when 1 concurrent operations attempts to read a variable while at some undetermined time another concurrent operations is attempting to write to the same variable

```golang
    /*
        There will be 3 possible outcomes:
        - Nothing is printed => data++ happen before if the condition
        - Value 0 is printed => the if condition and print happen before  data++
        - Value 1 is printed => the if condition happen but data++ before print
    */
    var data int
    go func() {
        data++
    }
    if data == 0 {
        fmt.Printf("The value is %v.\n", data)
    }
```

=> We need to iterate through the possible scenarios.

### Atomicity

- `context`. Something may be atomic in one context but not another.
