package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var input, banner string
	argsCount := len(os.Args)

	// 1. Handle command-line arguments or prompt the user
	if argsCount >= 3 {
		// If user provides text and banner file
		input = os.Args[1]
		banner = os.Args[2]
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		// If text is missing
		if argsCount < 2 {
			fmt.Print("Enter the text to render: ")
			scanner.Scan()
			input = scanner.Text()
		} else {
			input = os.Args[1]
		}

		// If banner is missing
		if argsCount < 3 {
			fmt.Print("Enter the banner file to use (e.g., shadow.txt): ")
			scanner.Scan()
			banner = scanner.Text()
		} else {
			banner = os.Args[2]
		}
	}

	// 2. Replace \n literals with real newlines
	if input == `\n` {
		fmt.Println()
		return
	}
	if input == "" {
		return
	}
	input = strings.ReplaceAll(input, `\n`, "\n")

	// 3. Open the banner file provided by user
	file, err := os.Open(banner)
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}
	defer file.Close()

	// 4. Read banner file into mapping
	scanner := bufio.NewScanner(file)
	bannerMap := getBannerMapping(scanner)
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error : %s", err)
	}

	// 5. Split input by \n but retain newlines
	sSlice := splitStrByNewLines(input)

	// 6. Render ASCII art
	resStr := getResultAscii(sSlice, bannerMap)

	// 7. Print the result
	fmt.Printf("%v", resStr)
}
