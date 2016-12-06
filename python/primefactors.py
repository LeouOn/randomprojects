def factors(n):
	factors = []
	for i in range(2,n):
		if n%i == 0:
			factors.append([i])
	return factors

def isPrime(n):
	primes = []
	for i in n:
		if (n % i) == 0:
			break
	else:
		primes.append([i])
	return primes
def Test(n):
	for i in n:
		print i

def primeFactors(n):
	return isPrime(factors(n))

print primeFactors(600851475143)

