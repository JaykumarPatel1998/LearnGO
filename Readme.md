## Hey Folks!!

This repository is me deep diving & learning Go Lang


# Learnings

1. **Array Initialization**: In Go, when you define an array with a fixed size, for example:
   
   ```go
   var arr [5]int
   ```

   This allocation happens at compile time because the size of the array is a constant integer, known at compile time. The compiler can determine how much memory to allocate for the array based on its fixed size. Arrays in Go have a fixed size that cannot change, and this is why the memory allocation is handled at compile time.

2. **Slices and the `make` Function**: When you use the `make()` function to initialize a slice, like this:

   ```go
   s := make([]int, 5)
   ```

   The memory allocation occurs at **runtime**. Slices in Go are more flexible than arrays; they have a dynamically resizable length and capacity, so their exact memory requirements are determined when the program runs. This flexibility means that you do not need to specify a compile-time constant for the length or capacity of a slice. Instead, you provide the desired length and capacity to `make()` at runtime.

   It's important to note, however, that the length and capacity **must still be integer values** (usually literals or computed values) when calling `make()`. Go's type system does not allow arbitrary non-integer types for specifying slice length or capacity, but these values can come from variables that are only known at runtime.

In summary:
- Arrays require a fixed integer length at compile time, and their memory allocation is handled at compile time.
- Slices, created with `make`, have their length and capacity determined at runtime, allowing for dynamic sizing and runtime memory allocation. However, the `make` function still requires an integer length and capacity as input parameters, even if those integers are determined at runtime.

