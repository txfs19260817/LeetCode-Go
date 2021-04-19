package _204_Count_Primes

func countPrimes(n int) int {
	ans, isPrime := 0, make([]bool, n) // an array of Bool values, indexed by 2 to n,
	for i := 2; i < n; i++ {           // initially [2:n) set to true
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ { // composite numbers range
		if isPrime[i] { // count up from the smallest marked prime
			for j := i * i; j < n; j += i { // j = i^2, i^2+i, i^2+2i, ..., < n
				isPrime[j] = false
			}
		}
	}
	for _, n := range isPrime { // count prime numbers
		if n == true {
			ans++
		}
	}
	return ans
}
