package main

import "testing"

func TestFibonacci(t *testing.T) {
	if Fibonacci(0) != 0 {
		t.Errorf("Expected %v, Actual %v", 0, Fibonacci(0))
	}
	if Fibonacci(1) != 1 {
		t.Errorf("Expected %v, Actual %v", 1, Fibonacci(1))
	}
	if Fibonacci(6) != 8 {
		t.Errorf("Expected %v, Actual %v", 8, Fibonacci(6))
	}
	if Fibonacci(21) != -1 {
		t.Errorf("Expected %v, Actual %v", -1, Fibonacci(21))
	}
	if Fibonacci(-1) != -1 {
		t.Errorf("Expected %v, Actual %v", -1, Fibonacci(-1))
	}
}
