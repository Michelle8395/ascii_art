package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	asciiStart = 32
	asciiEnd   = 126
	charHeight = 8
)

// Split input string by newline
func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

// Read banner file into slice of lines
func readBannerFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Build map[rune][]string where each rune maps to its ASCII art
func buildBannerMap(lines []string) map[rune][]string {
	banner := make(map[rune][]string)
	index := 0

	for ch := rune(asciiStart); ch <= asciiEnd; ch++ {
		banner[ch] = lines[index : index+charHeight]
		index += charHeight + 1 // skip empty line
	}

	return banner
}

// Print one word in ASCII
func printWord(word string, banner map[rune][]string) {
	for row := 0; row < charHeight; row++ {
		for _, ch := range word {
			if art, ok := banner[ch]; ok {
				fmt.Print(art[row])
			}
		}
		fmt.Println()
	}
}

// Print full ASCII art (handles newlines)
func printASCII(lines []string, banner map[rune][]string) {
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		printWord(line, banner)
	}
}

func main() {
	// Handle missing argument
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	lines := splitLines(input)

	// Read banner file
	bannerLines, err := readBannerFile("banner/standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Build ASCII map
	bannerMap := buildBannerMap(bannerLines)

	// Print result
	printASCII(lines, bannerMap)
}
