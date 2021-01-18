package sort

func Insert(b []byte) []byte {
	if len(b) < 2 {
		return b
	}

	for i := 1; i < len(b); i++ {
		for j := i; j > 0 && b[j-1] > b[j]; j-- {
			b[j-1], b[j] = b[j], b[j-1]
		}
	}

	return b
}
