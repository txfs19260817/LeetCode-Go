package k

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getUserMinMax(t *testing.T) {
	logs1 := [][]string{
		{"58523", "user_1", "resource_1"},
		{"62314", "user_2", "resource_2"},
		{"54001", "user_1", "resource_3"},
		{"200", "user_6", "resource_5"},
		{"215", "user_6", "resource_4"},
		{"54060", "user_2", "resource_3"},
		{"53760", "user_3", "resource_3"},
		{"58522", "user_22", "resource_1"},
		{"53651", "user_5", "resource_3"},
		{"2", "user_6", "resource_1"},
		{"100", "user_6", "resource_6"},
		{"400", "user_7", "resource_2"},
		{"100", "user_8", "resource_6"},
		{"54359", "user_1", "resource_3"},
	}

	expected1 := map[string][]int{
		"user_1":  {54001, 58523},
		"user_2":  {54060, 62314},
		"user_3":  {53760, 53760},
		"user_5":  {53651, 53651},
		"user_6":  {2, 215},
		"user_7":  {400, 400},
		"user_8":  {100, 100},
		"user_22": {58522, 58522},
	}

	result1 := getUserMinMax(logs1)
	assert.Equal(t, expected1, result1)
}

func Test_mostRequestedResource(t *testing.T) {
	logs1 := [][]string{
		{"58523", "user_1", "resource_1"},
		{"62314", "user_2", "resource_2"},
		{"54001", "user_1", "resource_3"},
		{"200", "user_6", "resource_5"},
		{"215", "user_6", "resource_4"},
		{"54060", "user_2", "resource_3"},
		{"53760", "user_3", "resource_3"},
		{"58522", "user_22", "resource_1"},
		{"53651", "user_5", "resource_3"},
		{"2", "user_6", "resource_1"},
		{"100", "user_6", "resource_6"},
		{"400", "user_7", "resource_2"},
		{"100", "user_8", "resource_6"},
		{"54359", "user_1", "resource_3"},
	}

	res, count := mostRequestedResource(logs1)
	assert.Equal(t, "resource_3", res)
	assert.Equal(t, 3, count)
}

func Test_transitionGraph(t *testing.T) {
	logs1 := [][]string{
		{"58523", "user_1", "resource_1"},
		{"62314", "user_2", "resource_2"},
		{"54001", "user_1", "resource_3"},
		{"200", "user_6", "resource_5"},
		{"215", "user_6", "resource_4"},
		{"54060", "user_2", "resource_3"},
		{"53760", "user_3", "resource_3"},
		{"58522", "user_22", "resource_1"},
		{"53651", "user_5", "resource_3"},
		{"2", "user_6", "resource_1"},
		{"100", "user_6", "resource_6"},
		{"400", "user_7", "resource_2"},
		{"100", "user_8", "resource_6"},
		{"54359", "user_1", "resource_3"},
	}

	expected1 := map[string]map[string]float64{
		"START": {
			"resource_1": 0.25,
			"resource_2": 0.125,
			"resource_3": 0.5,
			"resource_6": 0.125,
		},
		"resource_1": {
			"resource_6": 0.333,
			"END":        0.667,
		},
		"resource_2": {
			"END": 1.0,
		},
		"resource_3": {
			"END":        0.4,
			"resource_1": 0.2,
			"resource_2": 0.2,
			"resource_3": 0.2,
		},
		"resource_4": {
			"END": 1.0,
		},
		"resource_5": {
			"resource_4": 1.0,
		},
		"resource_6": {
			"END":        0.5,
			"resource_5": 0.5,
		},
	}

	result1 := transitionGraph(logs1)

	// Compare maps with float tolerance
	// Since we rounded to 3 decimal places in implementation, exact match might be tricky with float math
	// but the implementation explicitly rounds.
	// However, map equality in Go with floats can be strict.
	// Let's use a helper or just check if the values are close enough.
	// Or since we rounded in the impl, we can try direct equality.

	if !reflect.DeepEqual(expected1, result1) {
		// Debug print
		fmt.Println("Expected:", expected1)
		fmt.Println("Got:", result1)
		t.Errorf("transitionGraph() did not match expected")
	}

	logs2 := [][]string{
		{"300", "user_1", "resource_3"},
		{"599", "user_1", "resource_3"},
		{"900", "user_1", "resource_3"},
		{"1199", "user_1", "resource_3"},
		{"1200", "user_1", "resource_3"},
		{"1201", "user_1", "resource_3"},
		{"1202", "user_1", "resource_3"},
	}

	expected2 := map[string]map[string]float64{
		"START":      {"resource_3": 1.0},
		"resource_3": {"resource_3": 0.857, "END": 0.143},
	}

	result2 := transitionGraph(logs2)
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("transitionGraph() logs2 did not match expected")
	}
}
