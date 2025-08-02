package main

import (
	"testing"
)

func TestGenerateRandomElementsNormal(t *testing.T) {
	tests := []int{2, 3, 10000, 8, SIZE, SIZE - 1}

	for _, test := range tests {
		data := generateRandomElements(test)
		if len(data) != test {
			t.Errorf("Expected leight %d, got %d", test, len(data))
		}
	}
}

func TestGenerateRandomElementsInvalid(t *testing.T) {
	tests := []int{-1, 0, 1}

	for _, test := range tests {
		data := generateRandomElements(test)

		if data != nil {
			t.Errorf("error: data not nil")
		}
	}
}

func TestMaximumNormal(t *testing.T) {
	tests := []struct {
		data []int
		max  int
	}{
		{[]int{2, 3, 78, 2, 7, 99, 98, 801}, 801},
		{[]int{1, 110}, 110},
		{[]int{2, 0, 4}, 4},
		{[]int{5, 0}, 5},
	}

	for _, test := range tests {
		countMax := maximum(test.data)

		if test.max != countMax {
			t.Errorf("error: count not maximum")
		}
	}

}

func TestMaximumEmpty(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Error("error: slice not 0")
	}
}

// Пишите тесты в этом файле
