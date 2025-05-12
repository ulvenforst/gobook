package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	_ "github.com/ulvenforst/gobook/cmd/chap3"
	_ "github.com/ulvenforst/gobook/cmd/chap3/exercises"

	// "github.com/ulvenforst/gobook/cmd/chap4"
	"github.com/ulvenforst/gobook/cmd/chap4/exercises"
	_ "github.com/ulvenforst/gobook/cmd/chap4/exercises"
)

func main() {
	// chap3.RunSurface()
	// chap3.MandelbrotPlot()
	// str1 := "lll"
	// str2 := "lLl"
	// fmt.Println(exercises.AreAnagrams(str1, str2))

	// fmt.Println(exercises.RunCompareSHA256("x", "X"))

	// test := make([]int, 5, 10)
	//
	// a := append(test, 6)
	//
	// test[1] = 9
	//
	// fmt.Println(a)
	// fmt.Println(test)

	// chap4.LearningMaps()
	// exercises.RunPrintSHA()

	// s := []int{0, 1, 2, 3, 4, 5}
	//
	// if err := exercises.RotateSinglePass(s, 4); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(s)

	// s := []string{"Hello", "world", "world", "we", "are", "are", "are", "learning", "go", "go", "go"}
	//
	// fmt.Println(s)
	// s = exercises.EliminateAdj(s)
	// fmt.Println(s)

	// Exercise 4.9

	// counts := make(map[string]int)
	//
	// const PATH string = "cmd/experiments/textFiles/helloWorld.txt"
	// f, err := os.Open(PATH)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	//
	// exercises.WordFreq(f, counts)
	//
	// for key, value := range counts {
	// 	fmt.Printf("%8s\t%d\n", key, value)
	// }
	//

	// fmt.Println(exercises.Comma("11111123"))

	// use `go run cmd/main.go cmd/env.go` to execute it
	fmt.Printf("Enter a name of a movie to download its image: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	if err := exercises.Poster(input.Text()); err != nil {
		log.Fatal(err)
	}
}
