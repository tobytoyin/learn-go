package myModule

import "myProject/myModule"
import "testing"

// more advance testing can involes
// using `struct` to create inputs-result pairs
type testpair struct { 
	values []float64
	expectedResult float64
}

// create test scenarios
var testScenarios = []testpair{
	// { inputs, expected result }
	{ []float64{1, 2}, 1.5 }, 
	{ []float64{2, 2}, 2.0 }, 
	{ []float64{-1, 1}, 0.0 }, 
}

// writing for-loops to test all the scenarios
func TestScenariosAverage(t *testing.T) {
	for _, pair := range testScenarios {
		v := myModule.Average(pair.values)
		if v != pair.expectedResult {
			t.Error("test result is not matched")
		}
	}
}
