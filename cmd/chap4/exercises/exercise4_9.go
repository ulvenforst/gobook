package exercises

import (
	"bufio"
	"os"
	"strings"
)

// Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.

// WordFreq reports the frequency of each word in an input text file.
func WordFreq(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[strings.ToLower(input.Text())]++
	}
}
