package utils

import (
	"reflect"
	"testing"
)

// 测试 Min 函数
func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test with normal integers",
			input:    []int{3, 1, 4, 2},
			expected: 1,
		},
		{
			name:     "Test with all same integers",
			input:    []int{5, 5, 5},
			expected: 5,
		},
		{
			name:     "Test with single integer",
			input:    []int{7},
			expected: 7,
		},
		{
			name:     "Test with empty input",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Min(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// 测试 Max 函数
func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test with normal integers",
			input:    []int{3, 1, 4, 2},
			expected: 4,
		},
		{
			name:     "Test with all same integers",
			input:    []int{5, 5, 5},
			expected: 5,
		},
		{
			name:     "Test with single integer",
			input:    []int{7},
			expected: 7,
		},
		{
			name:     "Test with empty input",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Max(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
