package chap2

// Pc[i] is the population count of i.
var Pc [256]byte

// Pc is a lookup table for the population count of 8-bit integers.
func init() {
	for i := range Pc {
		Pc[i] = Pc[i/2] + byte(i&1)
	}
}
