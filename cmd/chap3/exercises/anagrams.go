package exercises

import (
	"slices"
	"strings"
)

// Declarative solution: O(n log n)
func AreAnagrams(str1, str2 string) bool {
	str1Runes := []rune(strings.ToLower(str1))
	slices.Sort(str1Runes)

	str2Runes := []rune(strings.ToLower(str2))
	slices.Sort(str2Runes)
	return slices.Equal(str1Runes, str2Runes)
}

//// Imperative solution: O(n)
// func AreAnagrams(str1, str2 string) bool {
// 	counts := map[rune]int{}
//
// 	for _, r := range str1 {
// 		counts[r]++
// 	}
// 	for _, r := range str2 {
// 		counts[r]--
// 	}
// 	for _, c := range counts {
// 		if 0 != c {
// 			return false
// 		}
// 	}
// 	return true
// }
