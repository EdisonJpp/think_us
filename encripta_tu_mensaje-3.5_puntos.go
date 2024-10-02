package main

import (
	"fmt"
	"strings"
)

func x(key, message string) string {
	if message == "" {
		return ""
	}
	if key == "" {
		key = "DCJ"
	}

	vowels := "aeiouAEIOU"
	var result strings.Builder

	for _, char := range message {
		if strings.ContainsRune(vowels, char) {
			result.WriteString(key)
		}
		result.WriteRune(char)
	}
	return result.String()
}

func main() {
	key := "dcj"
	message := "I love prOgrAmming!"
	encrypted := x(key, message)
	fmt.Println(encrypted)
}
