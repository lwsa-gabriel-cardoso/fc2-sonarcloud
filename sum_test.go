package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 2)

	if result != 4 {
		t.Errorf("Sum(2, 2) = %d; want 4", result)
	}

	result = sub(2, 2)
	if result != 0 {
		t.Errorf("Sub(2, 2) = %d; want 0", result)
	}

	result = times(2, 2)
	if result != 4 {
		t.Errorf("Times(2, 2) = %d; want 4", result)
	}

	result = sumX(2, 2)
	if result != 6 {
		t.Errorf("SumX(2, 2) = %d; want 6", result)
	}

}
