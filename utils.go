package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ----------------- Helper functions -----------------

func getBannerMapping(scanner *bufio.Scanner) map[rune][]string {
	i := 0
	curRune := ' '
	bannerMap := map[rune][]string{
		' ': {},
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		bannerMap[curRune] = append(bannerMap[curRune], line)
		if i > 0 && (i+1)%8 == 0 {
			curRune++
			if i <= 126 {
				bannerMap[curRune] = []string{}
			}
		}
		if len(line) > 0 {
			i++
		}
	}

	return bannerMap
}

func getAsciiLine(s string, bannerMap map[rune][]string) string {
	var lineRes strings.Builder
	for i := 0; i < 8; i++ {
		for _, r := range s {
			group, exists := bannerMap[r]
			if exists {
				lineRes.WriteString(group[i])
			}
		}
		if i != 7 {
			lineRes.WriteString("\n")
		}
	}
	return lineRes.String()
}

func getResultAscii(sSlice []string, bannerMap map[rune][]string) string {
	var resStr strings.Builder
	for _, s := range sSlice {
		sR := []rune(s)
		if len(sR) == 1 && sR[0] == '\n' {
			resStr.WriteString("\n")
			continue
		}
		resStr.WriteString(getAsciiLine(s, bannerMap))
	}
	resStr.WriteString("\n")
	return resStr.String()
}

func splitStrByNewLines(s string) []string {
	var tokens []string
	current := ""

	for _, r := range s {
		if r == '\n' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, "\n")
		} else {
			current += string(r)
		}
	}

	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}

// ----------------- New RenderAscii function -----------------

// RenderAscii renders the input string using the given banner file.
// Returns the ASCII art as a string or an error if something goes wrong.
func RenderAscii(input string, bannerFile string) (string, error) {
	// 1. Open banner file
	file, err := os.Open(bannerFile)
	if err != nil {
		return "", fmt.Errorf("failed to open banner file: %w", err)
	}
	defer file.Close()

	// 2. Build banner mapping
	scanner := bufio.NewScanner(file)
	bannerMap := getBannerMapping(scanner)
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	// 3. Split input by \n but retain newlines
	sSlice := splitStrByNewLines(input)

	// 4. Generate ASCII
	resStr := getResultAscii(sSlice, bannerMap)

	return resStr, nil
}
