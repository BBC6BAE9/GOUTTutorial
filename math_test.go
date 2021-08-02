package main

import "testing"

func TestMyMod(t *testing.T) {

	tests := []struct {
		a int
		b int
		want int
		desc string
	}{
		{1, 2, 1, "a小于b"},
		{1, 2, 1, "a小于b"},
		{1, 2, 1, "a小于b"},
		{1, 2, 1, "a小于b"},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := MyMod(tt.a, tt.b); got != tt.want {
				t.Errorf("MyMod() = %v, want %v", got, tt.want)
			}
		})
	}
}
