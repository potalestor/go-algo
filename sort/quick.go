package sort

func Quick(b []byte) []byte {
	if len(b) < 2 {
		return b
	}

	quick(b, 0, len(b)-1)

	return b
}

func quick(b []byte, f, e int) {
	if f >= e {
		return
	}

	l, r := f, e

	p := b[(l+r)>>1]

	for l <= r {
		for b[l] < p {
			l++
		}

		for b[r] > p {
			r--
		}

		if l <= r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}
	quick(b, f, r)
	quick(b, l, e)
}

// func Quick(b []byte) []byte {
// 	if len(b) < 2 {
// 		return b
// 	}

// 	l, r := 0, len(b)-1

// 	pivot := b[r]

// 	// i - итерация по массиву
// 	// l - индекс для опорного элемента после разделения
// 	for i := range b {
// 		if b[i] < pivot {
// 			b[i], b[l] = b[l], b[i]
// 			l++
// 		}
// 	}

// 	b[r], b[l] = b[l], b[r]

// 	Quick(b[:l])
// 	Quick(b[l+1:])

// 	return b
// }
