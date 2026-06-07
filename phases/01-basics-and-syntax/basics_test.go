package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	if AppName != "TodoApp" {
		t.Errorf("Expected AppName to be 'TodoApp', got '%s'", AppName)
	}
}

func TestFormatTodo(t *testing.T) {
	tests := []struct {
		title     string
		completed bool
		expected  string
	}{
		{"Buy Milk", true, "Buy Milk [Completed]"},
		{"Read Book", false, "Read Book [Pending]"},
		{"", true, " [Completed]"},
	}

	for _, tt := range tests {
		result := FormatTodo(tt.title, tt.completed)
		if result != tt.expected {
			t.Errorf("FormatTodo(%q, %t) = %q; expected %q", tt.title, tt.completed, result, tt.expected)
		}
	}
}

func TestIsPriorityHigh(t *testing.T) {
	tests := []struct {
		priority int
		expected bool
	}{
		{5, true},
		{4, false},
		{6, true},
		{0, false},
		{-1, false},
	}

	for _, tt := range tests {
		result := IsPriorityHigh(tt.priority)
		if result != tt.expected {
			t.Errorf("IsPriorityHigh(%d) = %t; expected %t", tt.priority, result, tt.expected)
		}
	}
}

func TestCalculateCompletionRate(t *testing.T) {
	tests := []struct {
		completed int
		total     int
		expected  float64
	}{
		{3, 10, 0.3},
		{0, 5, 0.0},
		{5, 0, 0.0}, // Division by zero check
		{5, 5, 1.0},
	}

	for _, tt := range tests {
		result := CalculateCompletionRate(tt.completed, tt.total)
		if result != tt.expected {
			t.Errorf("CalculateCompletionRate(%d, %d) = %v; expected %v", tt.completed, tt.total, result, tt.expected)
		}
	}
}
