package main

import (
	"bufio"
	"strings"
)

func getBannerMapping(scanner *bufio.Scanner) map[rune][]string {
	i := 0
	curRune := ' '
	bannerMap := map[rune][]string{
		' ': {},
	}

	// Iterate over the scanner to read line by line
	for scanner.Scan() {
		// Get the current line as a string (newline termination is stripped by default)
		line := scanner.Text()
		if len(line) == 0 { // If line is empty ignore it
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

// Returns a string containing the 8 layers of a string that can then be printed to output
func getAsciiLine(s string, bannerMap map[rune][]string) string {
	var lineRes strings.Builder
	for i := 0; i < 8; i++ { // Loop 8 times
		for _, r := range s {
			group, exists := bannerMap[r]
			if exists {
				lineRes.WriteString(group[i])
			}
		}
		if i != 7 { // Only
			lineRes.WriteString("\n")
		}
	}

	return lineRes.String()
}

// For each string in sSlice, its expected representation in ascii is obtained & appended to a result string & then returned
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

// Split input using \n but still retain the \n
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
