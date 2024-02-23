# Golang Data Types

## Numbers

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
