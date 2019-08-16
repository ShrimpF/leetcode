package algo

func memorize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if _, ok := cache[n]; !ok {
			cache[n] = f(n)
		}
		return cache[n]
	}
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-2) + fibo(n-1)
}
