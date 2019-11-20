package loop

func FibonacciUsingLoops(x int) int {
    fib := make([]int, x+1, x+2)
    if x < 2 {
        fib = fib[0:2]
    }
    fib[0] = 0
    fib[1] = 1
    for i := 2; i <= x; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    return fib[x]
}



