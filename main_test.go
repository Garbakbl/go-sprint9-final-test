package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name   string
		size   int
		resLen int
	}{
		{name: "empty", size: 0, resLen: 0},
		{name: "negative", size: -10, resLen: 0},
		{name: "positive", size: 10, resLen: 10},
		{name: "bigPositive", size: 999999999, resLen: 999999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.resLen, len(generateRandomElements(tt.size)))
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		res  int
	}{
		{name: "only", arr: []int{1000}, res: 1000},
		{name: "normal", arr: []int{1, 2, 5, 100500}, res: 100500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.res, maximum(tt.arr))
		})
	}
}
