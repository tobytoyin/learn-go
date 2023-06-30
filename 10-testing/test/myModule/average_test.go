package myModule

import "myProject/myModule"
import "testing"

func TestAverage(t *testing.T) {
	v := myModule.Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

