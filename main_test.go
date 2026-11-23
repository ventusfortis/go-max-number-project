package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "negative size", size: -1},
		{name: "zero size", size: 0},
		{name: "one element", size: 1},
		{name: "many elements", size: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != max(tt.size, 0) {
				t.Fatalf("expected length %d, got %d", max(tt.size, 0), len(got))
			}
			for i, value := range got {
				if value <= 0 {
					t.Errorf("element %d must be positive, got %d", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     int
	}{
		{name: "nil slice", elements: nil, want: 0},
		{name: "empty slice", elements: []int{}, want: 0},
		{name: "one element", elements: []int{7}, want: 7},
		{name: "maximum at beginning", elements: []int{9, 2, 3}, want: 9},
		{name: "maximum at end", elements: []int{1, 2, 10}, want: 10},
		{name: "maximum in middle", elements: []int{1, 12, 3}, want: 12},
		{name: "equal elements", elements: []int{5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.elements); got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     int
	}{
		{name: "nil slice", elements: nil, want: 0},
		{name: "empty slice", elements: []int{}, want: 0},
		{name: "one element", elements: []int{7}, want: 7},
		{name: "fewer than eight elements", elements: []int{2, 9, 1}, want: 9},
		{name: "exactly eight elements", elements: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "remainder is included", elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
		{name: "maximum in first chunk", elements: []int{100, 2, 3, 4, 5, 6, 7, 8, 9}, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.elements); got != tt.want {
				t.Errorf("maxChunks() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	elements := generateRandomElements(10_003)
	if got, want := maxChunks(elements), maximum(elements); got != want {
		t.Fatalf("maxChunks() = %d, maximum() = %d", got, want)
	}
}
