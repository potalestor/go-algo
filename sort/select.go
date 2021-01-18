package sort

func Select(b []byte) []byte {
	if len(b) < 2 {
		return b
	}

	for i := 0; i < len(b); i++ {
		j := min(b, i)
		b[j], b[i] = b[i], b[j]
	}

	return b
}

func min(b []byte, start int) int {
	m := start

	for i := start; i < len(b); i++ {
		if b[i] < b[m] {
			m = i
		}
	}

	return m
}
