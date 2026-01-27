package main

import (
	"os"
	"reflect"
	"testing"
)

// ---------- splitLines tests ----------

func TestSplitLinesSingleLine(t *testing.T) {
	input := "Hello"
	expected := []string{"Hello"}

	result := splitLines(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSplitLinesMultipleLines(t *testing.T) {
	input := "Hello\nThere"
	expected := []string{"Hello", "There"}

	result := splitLines(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSplitLinesEmptyLine(t *testing.T) {
	input := "Hello\n\nThere"
	expected := []string{"Hello", "", "There"}

	result := splitLines(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// ---------- banner file tests ----------

func TestReadBannerFile(t *testing.T) {
	_, err := readBannerFile("banner/standard.txt")
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}
}

func TestBuildBannerMapSize(t *testing.T) {
	lines, err := readBannerFile("banner/standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	banner := buildBannerMap(lines)

	if len(banner) != 95 {
		t.Errorf("Expected 95 ASCII characters, got %d", len(banner))
	}
}

func TestBuildBannerMapCharacterHeight(t *testing.T) {
	lines, err := readBannerFile("banner/standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	banner := buildBannerMap(lines)

	if len(banner['A']) != charHeight {
		t.Errorf("Expected height %d, got %d", charHeight, len(banner['A']))
	}
}

// ---------- safety tests ----------

func TestInvalidBannerPath(t *testing.T) {
	_, err := readBannerFile("banner/does_not_exist.txt")
	if err == nil {
		t.Errorf("Expected error for missing banner file")
	}
}

func TestMainNoArgs(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd"}
	main() // should NOT panic
}
