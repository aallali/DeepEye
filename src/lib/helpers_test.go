package lib

import (
	"reflect"
	"testing"

	"github.com/dlclark/regexp2"
)

func TestIfElse(t *testing.T) {
	// Test case 1: When cond is true
	result := ifElse(true, "true value", "false value")
	expected := "true value"
	if result != expected {
		t.Errorf("ifElse(true, \"true value\", \"false value\") = %v; expected %v", result, expected)
	}

	// Test case 2: When cond is false
	result = ifElse(false, "true value", "false value")
	expected = "false value"
	if result != expected {
		t.Errorf("ifElse(false, \"true value\", \"false value\") = %v; expected %v", result, expected)
	}

	// Test case 3: When cond is true and vTrue is an integer
	result = ifElse(true, 10, 20)
	num_expected := 10
	if result != num_expected {
		t.Errorf("ifElse(true, 10, 20) = %v; expected %v", result, expected)
	}

	// Test case 4: When cond is false and vFalse is a float
	result = ifElse(false, "string value", 3.14)
	float_expected := 3.14
	if result != float_expected {
		t.Errorf("ifElse(false, \"string value\", 3.14) = %v; expected %v", result, expected)
	}
}

func TestRegexp2FindAllString(t *testing.T) {
	// Test case 1: No matches
	re := regexp2.MustCompile("abc", 0)
	s := "defg"
	expected := []string{}
	result := regexp2FindAllString(re, s)
	if len(result) != len(expected) {
		t.Errorf("regexp2FindAllString(re, s) returned %v; expected %v", result, expected)
	}

	// Test case 2: Single match
	re = regexp2.MustCompile("[0-9]+", 0)
	s = "The number is 123."
	expected = []string{"123"}
	result = regexp2FindAllString(re, s)
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("regexp2FindAllString(re, s) returned %v; expected %v", result, expected)
	}

	// Test case 3: Multiple matches
	re = regexp2.MustCompile("foo", 0)
	s = "foo foo foo"
	expected = []string{"foo", "foo", "foo"}
	result = regexp2FindAllString(re, s)
	if len(result) != len(expected) {
		t.Errorf("regexp2FindAllString(re, s) returned %v; expected %v", result, expected)
	}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("regexp2FindAllString(re, s) returned %v; expected %v", result, expected)
			break
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	// Test case 1: No duplicates
	input := []string{"a", "b", "c"}
	expected := []string{"a", "b", "c"}
	result := removeDuplicates(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeDuplicates(%v) returned %v; expected %v", input, result, expected)
	}

	// Test case 2: Duplicates within the input slice
	input = []string{"a", "b", "a", "c", "b", "a"}
	expected = []string{"a", "b", "c"}
	result = removeDuplicates(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeDuplicates(%v) returned %v; expected %v", input, result, expected)
	}

	// Test case 3: Empty string values and duplicates
	input = []string{"", "a", "", "b", "a", ""}
	expected = []string{"a", "b"}
	result = removeDuplicates(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeDuplicates(%v) returned %v; expected %v", input, result, expected)
	}

	// Test case 4: Empty input slice
	input = []string{}
	expected = []string{}
	result = removeDuplicates(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeDuplicates(%v) returned %v; expected %v", input, result, expected)
	}
}
