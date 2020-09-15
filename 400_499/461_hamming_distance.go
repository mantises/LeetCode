package _400_499

func hammingDistance(x int, y int) int {
	xor := x ^ y
	count := 0
	for xor != 0 {
		count++
		xor = xor & (xor - 1)
	}
	return count
}
