package main

import "testing"

func FuzzPositiveNum(f *testing.F) {
	seedNum := 90
	f.Add(seedNum)
	f.Fuzz(func(t *testing.T, a int) { // fuzz target
		if a != PositiveNum(a) {
			t.Fail()
		}
	})
}
