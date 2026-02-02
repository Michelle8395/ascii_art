package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	if input == `\n` {
		fmt.Println()
		return
	}

	if input == "" {
		return
	}

	input = strings.ReplaceAll(os.Args[1], `\n`, "\n")

	// 1. Open the file
	file, err := os.Open("shadow.txt")
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}

	// 2. Defer closing the file until the main function returns
	defer file.Close()

	// 3. Create a new scanner object for the file
	scanner := bufio.NewScanner(file)

	// 4. Create mapping for {[rune] : ascii_string_layers_slice } as map[rune][]string
	bannerMap := getBannerMapping(scanner)

	// 5. Check for errors that occurred during scanning (EOF is not an error)
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error : %s", err)
	}

	// 6. Split input using \n but still retain the \n
	sSlice := splitStrByNewLines(input)

	// 7. Get result string from each of those split strings
	resStr := getResultAscii(sSlice, bannerMap)

	// 8. Print the result
	fmt.Printf("%v", resStr)
}
