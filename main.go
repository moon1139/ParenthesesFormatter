package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	inputFile, err := os.Open("input.txt") // Open the input file
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.txt") // Create the output file
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		expression := scanner.Text()
		result := formatParentheses(expression)
		
		outputConsole := fmt.Sprintf("%s => %s\n", expression, result)
		outputLine := fmt.Sprintf("%s\n", result)
		fmt.Print(outputConsole) 	   // Print to console
		writer.WriteString(outputLine) // Write to output file
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}

	writer.Flush() // Ensure all data is written to the file
}
