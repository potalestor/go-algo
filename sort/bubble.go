package sort

func Bubble(b []byte) []byte {
	if len(b) < 2 {
		return b
	}

	for i := 0; i < len(b); i++ {
		for j := 1; j < len(b)-i; j++ {
			if b[j-1] > b[j] {
				b[j-1], b[j] = b[j], b[j-1]
			}
		}
	}

	return b
}
