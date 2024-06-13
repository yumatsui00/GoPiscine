package piscine

func fibo(i int, j int, index int) int {
	ans := i + j
	if index == 0 {
		return ans
	}
	return fibo(j, ans, index - 1)
}

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	return fibo(0, 1, index - 2)
}