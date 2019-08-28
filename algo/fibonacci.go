package algo

// Memorize is memorize the result with map. this prevents from recalculation
func Memorize(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = Memorize(n-2, memo) + Memorize(n-1, memo)
	}
	return memo[n]
}

// Fibo return fibonacci number
func Fibo(n int) int {
	if n < 2 {
		return n
	}
	return Fibo(n-2) + Fibo(n-1)
}

//Fibo2 -- return fibonnacci num
// botton up solution
func Fibo2(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
