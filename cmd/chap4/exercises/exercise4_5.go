package exercises

// EliminateAdj removes adjacent duplicates from a slice of strings.
func EliminateAdj(s []string) []string {
	var write int
	for i := 1; i < len(s); i++ {
		if s[write] != s[i] {
			write++
			s[write] = s[i]
		}
	}
	return s[:write+1]
}
