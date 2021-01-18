package dp

// longest common substring problem is to find the longest string that is a substring of two strings.
func LongestCommonSubstring(s1, s2 string) string {
	var max byte

	index := 0
	tbl := newTable(s1, s2)

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				tbl.data[i+1][j+1] = tbl.data[i][j] + 1
				if max < tbl.data[i+1][j+1] {
					max = tbl.data[i+1][j+1]
					index = j + 1
				}
			}
		}
	}

	// fmt.Println(tbl)

	if max > 0 {
		return s2[index-int(max) : index]
	}

	return ""
}
