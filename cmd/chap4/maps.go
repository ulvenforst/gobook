package chap4

import "fmt"

func LearningMaps() {
	ages := map[string]int{
		"John": 25,
		"Jane": 30,
	}

	fmt.Println(ages["John"])
}
