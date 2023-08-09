package main

import (
	"testing"
)

func TestStockPicker(t *testing.T) {
	var tests = []struct {
		Name   string
		Sample []int
		Output int
	}{
		{"5", []int{10, 12, 4, 5, 9}, 5},
		{"8", []int{14, 20, 4, 12, 5, 11}, 8},
		{"16", []int{44, 30, 24, 32, 35, 30, 40, 38, 15}, 16},
		{"-1", []int{10, 9, 8, 2}, -1},
		{"0", []int{101, 79, 38, 24, 34}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			want := tt.Output
			got := StockPicker(tt.Sample)
			if want != got {
				t.Errorf("want %v got %v", want, got)
			}
		})
	}
}
