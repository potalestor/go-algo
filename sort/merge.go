package sort

func Merge(b []byte) []byte {
	if len(b) < 2 {
		return b
	}

	mid := len(b) >> 1

	l := b[:mid]
	r := b[mid:]

	return merge(Merge(l), Merge(r))
}

func merge(l, r []byte) []byte {
	size := len(l) + len(r)
	result := make([]byte, size)

	var li, ri int

	for i := 0; i < size; i++ {
		if li >= len(l) {
			result[i] = r[ri]
			ri++

			continue
		}

		if ri >= len(r) {
			result[i] = l[li]
			li++

			continue
		}

		if l[li] < r[ri] {
			result[i] = l[li]
			li++

			continue
		}

		result[i] = r[ri]
		ri++
	}

	return result
}
