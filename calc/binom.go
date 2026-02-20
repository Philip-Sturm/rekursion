package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	//Fall 0: k=n return 1
	if k == n || k == 0 {
		return 1
	}
	//Fall 1: k=1 return n
	if k == 1 {
		return n
	}

	return BinomialCoefficient(n-1, k) + BinomialCoefficient(n-1, k-1)
}
