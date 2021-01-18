package dp

import "testing"

func TestLongestCommonSubstring(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"abcdxyz", "xyzabcdxxx"}, "abcdx"},
		{"02", args{"a", "x"}, ""},
		{"03", args{"a", ""}, ""},
		{"04", args{"", "a"}, ""},
		{"05", args{"", ""}, ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubstring(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LongestCommonSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
