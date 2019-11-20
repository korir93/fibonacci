package recursion


func FibonacciUsingRecursion(x int) int {
    if x <= 1 {
        return x
    }
    return FibonacciUsingRecursion(x-1) + FibonacciUsingRecursion(x-2)
}
