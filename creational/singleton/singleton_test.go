package singleton

import "testing"

var counter1 *singleton
var expectedCounter *singleton
var currentCount int

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		// Test of acceptance criteria 1 failed
		t.Error("A new connection object must have been made")
	}
	expectedCounter = counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}
}

func TestIncrement(t *testing.T) {
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		// Test 2 Failed
		t.Error("Singleton instances must be different")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}