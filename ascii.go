package main

import (
	"fmt"
	"os"
	"strings"
)
func GenerateASCII(text, banner string) (string, error) {
	chars, err := loadBanner(banner)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	lines := strings.Split(text, "\\n")

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n") 
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if ch < 32 || ch > 126 {
					return "", fmt.Errorf("character %q is not supported", ch)
				}
				index := (int(ch)-32)*9 + 1 + row
				result.WriteString(chars[index])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
func loadBanner(name string) ([]string, error) {
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !validBanners[name] {
		return nil, fmt.Errorf("unknown banner %q — choose standard, shadow, or thinkertoy", name)
	}

	path := fmt.Sprintf("banner/%s.txt", name)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read banner file %q: %w", path, err)
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(content, "\n")
	return lines, nil
}
