package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want int
	}{
		{"Empty Input", []byte(""), 0},
		{"No Newlines", []byte("Hello World"), 0},
		{"Single Newline", []byte("Hello\nWorld"), 1},
		{"Multiple Lines with Ending Newline", []byte("Line1\nLine2\nLine3\n"), 3},
		{"Multiple Lines without Ending Newline", []byte("Line1\nLine2\nLine3"), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLines(tt.data); got != tt.want {
				t.Errorf("getLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want int
	}{
		{"Empty Input", []byte(""), 0},
		{"Single Word", []byte("Hello"), 1},
		{"Multiple Words", []byte("Hello World Go"), 3},
		{"Words with Newlines", []byte("Hello\nWorld Go"), 3},
		{"Mixed Whitespace", []byte("Hello   World\tGo\n"), 3},
		{"Only Spaces", []byte("   "), 0},
		{"Unicode Characters", []byte("こんにちは 世界"), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWords(tt.data); got != tt.want {
				t.Errorf("getWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChars(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want int
	}{
		{"Empty Input", []byte(""), 0},
		{"ASCII Characters", []byte("Hello"), 5},
		{"Unicode Characters", []byte("こんにちは"), 5},
		{"Mixed ASCII and Unicode", []byte("Hello世界"), 7},
		{"Emojis", []byte("😊😊"), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChars(tt.data); got != tt.want {
				t.Errorf("getChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want int
	}{
		{"Empty Input", []byte(""), 0},
		{"ASCII Characters", []byte("Hello"), 5},
		{"Unicode Characters", []byte("こんにちは"), 15},
		{"Mixed ASCII and Unicode", []byte("Hello 世界"), 12},
		{"Emojis", []byte("😊😊"), 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBytes(tt.data); got != tt.want {
				t.Errorf("getBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
