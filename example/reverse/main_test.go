package main

import "testing"

func TestReverse(t *testing.T) {
	// Given
	testcases := []struct {
		got, want string
	}{
		{"The quick brown fox jumped over the lazy dog", "god yzal eht revo depmuj xof nworb kciuq ehT"},
		{"", ""},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.got)
		if rev != tc.want {
			t.Errorf("got %q want %v", rev, tc.want)
		}
	}

}
