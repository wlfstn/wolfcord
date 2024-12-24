package wc

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ResourceLoadFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("file error: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("read error: %v", err)
	}

	output := strings.TrimSpace(string(content))
	return output
}

func ResourceLoadSQL(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	defer file.Close()

	query, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(query)
}
