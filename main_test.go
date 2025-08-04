package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected []int
	}{
		{
			name:     "Positive",
			size:     SIZE,
			expected: make([]int, SIZE),
		},
		{
			name:     "Zero",
			size:     0,
			expected: nil,
		},
		{
			name:     "Negative",
			size:     1,
			expected: nil,
		},
	}

	for _, test := range tests {
		result := generateRandomElements(test.size)
		if test.size > 1 {
			assert.Equal(t, test.size, len(result), test.name)
			assert.NotNil(t, result, test.name)
		} else {
			assert.Nil(t, result, test.name)
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Normal",
			data:     []int{3, 7, 2, 9, 5},
			expected: 9,
		},
		{
			name:     "Single",
			data:     []int{42},
			expected: 0,
		},
		{
			name:     "Empty slice",
			data:     []int{},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := maximum(test.data)
		assert.Equal(t, test.expected, result, test.name)
	}

}

// Пишите тесты в этом файле
