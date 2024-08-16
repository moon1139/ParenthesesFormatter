package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		expression := scanner.Text()
		result := formatParentheses(expression)
		
		// outputConsole := fmt.Sprintf("%s => %s\n", expression, result)
		// fmt.Print(outputConsole) 	   // Print to console
		outputLine := fmt.Sprintf("%s\n", result)
		writer.WriteString(outputLine)     // Write to stdout
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	writer.Flush() // Ensure all data is written to stdout
}
