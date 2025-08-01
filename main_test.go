package main

import (
	"testing"
)

func TestGenerateRandomElementsNormal(t *testing.T) {
	for i := 2; i < CHUNKS; i++ {
		data := generateRandomElements(i)
		if len(data) != i {
			t.Errorf("Expected leight %d, got %d", i, len(data))
		}
	}
}

func TestGenerateRandomElementsInvalid(t *testing.T) {
	for j := -1; j < 2; j++ {
		data := generateRandomElements(j)

		if data != nil {
			t.Errorf("error: data not nil")
		}
	}
}

func TestMaximumNormal(t *testing.T) {
	data := []int{3, 78, 2, 8, 99, 98, 801}
	max := 801

	countMax := maximum(data)
	if max != countMax {
		t.Errorf("error: count not maximum")
	}
}

func TestMaximumEmpty(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Error("error: slice not 0")
	}
}

// Пишите тесты в этом файле
