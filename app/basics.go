package main

import "fmt"

var AppName = "TodoApp"

func FormatTodo(title string, completed bool) string {
	// Hint: You can use standard string concatenation (+) or fmt.Sprintf
	taskStatus := "Pending"

	if completed {
		taskStatus = "Completed"
	}

	return fmt.Sprintf("%s [%s]", title, taskStatus)
}

func IsPriorityHigh(priority int) bool {
	if priority >= 5 {
		return true
	}
	return false
}

func CalculateCompletionRate(completedCount int, totalCount int) float64 {
	if totalCount == 0 {
		return 0.0
	}

	return float64(completedCount) / float64(totalCount)
}
