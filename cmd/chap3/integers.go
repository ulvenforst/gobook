package chap3

import "fmt"

func RunChap3() {
	var unsignedInt uint8 = 255
	var signedInt int8 = 127

	// It loops on its range when overflowing
	fmt.Println("Unsigned int:", unsignedInt+3)
	fmt.Println("Unsigned int:", signedInt+3)

	// DEALING WITH BITS
	var testByte byte = 1

	fmt.Printf("Unsigned int: %08b\n", testByte<<1)
	fmt.Printf("Unsigned int: %08b\n", ^testByte)

	// We often represent sets using bytes
	// Lets put the number 1 and 5 in a set
	var set1 uint8 = 1<<1 | 1<<5
	// Lets put the number 1 and 2 in a set
	var set2 uint8 = 1<<1 | 1<<2

	// Now we can perform bitwise operations on the sets
	// Union:
	fmt.Printf("set1 $\\cup$ set2: %08b\n", set1|set2)
	// Intersection:
	fmt.Printf("set1 $\\cap$ set2: %08b\n", set1&set2)
	// Difference:
	fmt.Printf("set1 $\\setminus$ set2: %08b\n", set1&^set2)
}
