package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Started Log Analyzer")

	// Check command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <logfile>")
		return
	}

	filename := os.Args[1]

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize counters
	totalLines := 0
	infoCount := 0
	errorCount := 0
	warningCount := 0

	// Map to store error message frequency
	errorMap := make(map[string]int)

	// Read file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		if strings.Contains(line, "INFO") {
			infoCount++
		}

		if strings.Contains(line, "WARNING") {
			warningCount++
		}

		if strings.Contains(line, "ERROR") {
			errorCount++

			parts := strings.SplitN(line, "ERROR", 2)
			if len(parts) == 2 {
				msg := strings.TrimSpace(parts[1])
				errorMap[msg]++
			}
		}
	}

	// Scanner error check
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Find most frequent error
	maxCount := 0
	maxMsg := ""

	for msg, count := range errorMap {
		if count > maxCount {
			maxCount = count
			maxMsg = msg
		}
	}

	// Print report
	fmt.Println("\nLog Analysis Report")
	fmt.Println("------------------")
	fmt.Println("Total Lines:", totalLines)
	fmt.Println("INFO:", infoCount)
	fmt.Println("WARNING:", warningCount)
	fmt.Println("ERROR:", errorCount)

	if maxCount > 0 {
		fmt.Printf("\nMost common ERROR:%s (%d times)\n", maxMsg, maxCount)
	} else {
		fmt.Println("\nNo ERROR logs found")
	}
}
