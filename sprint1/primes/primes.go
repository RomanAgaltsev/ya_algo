package primes

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func IsPrimeFast(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func IsPrimeTricky(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	return (num-1)%6 == 0 || (num+1)%6 == 0
}