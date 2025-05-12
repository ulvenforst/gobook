package exercises

import "errors"

func RotateLeft(s []int, n int) {
	Reverse(s[:n])
	Reverse(s[n:])
	Reverse(s)
}

func RotateSinglePass(s []int, n int) error {
	k := n
	lenS := len(s)
	if k <= lenS/2 {
		holder := make([]int, k)
		for i := range lenS {
			if i < k {
				holder[i] = s[lenS-n]
				s[lenS-n] = s[i]
				n--
			} else if i < lenS-k {
				s[i-k] = s[i]
			} else {
				s[i-k] = holder[n]
				n++
			}
		}
		return nil
	}
	return errors.New("Right now only k<=len(s)/2 is supported")
}

//// Solution by GPT, I didnt think bout copy
// func RotateSinglePass(s []int, n int) {
// 	n = n % len(s) // para evitar rotaciones mayores que el tamaño
// 	if n == 0 {
// 		return
// 	}
//
// 	// Copiar los primeros n elementos
// 	tmp := make([]int, n)
// 	copy(tmp, s[:n])
//
// 	// Mover los restantes a la izquierda
// 	copy(s, s[n:])
//
// 	// Copiar el buffer al final
// 	copy(s[len(s)-n:], tmp)
// }

//// What I'd use in real life
// func RotateSinglePass(s []int, n int) []int {
// 	return slices.Concat(s[n:], s[:n])
// }
