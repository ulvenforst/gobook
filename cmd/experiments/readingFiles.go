package experiments

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunReadFiles() {
	contentToWrite := "Este lo escribimos desde readingFiles.go\n"

	f, err := os.OpenFile("cmd/experiments/textFiles/helloWorld.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(contentToWrite)
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)

	input := bufio.NewScanner(f)

	for input.Scan() {
		fmt.Println(input.Text())
	}

}
