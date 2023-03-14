package area

// maxBound returns the maximum of two numbers.
func maxBound(b1, b2 float64) float64 {
	if b1 > b2 {
		return b1
	}
	return b2
}

// minBound returns the minimum of two numbers.
func minBound(b1, b2 float64) float64 {
	if b1 < b2 {
		return b1
	}
	return b2
}
