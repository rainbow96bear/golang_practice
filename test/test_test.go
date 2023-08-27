package main

import "testing"

func TestFunc (t *testing.T) {
	rst := practiceTest(2,3)
	if rst < 10 { 
		t.Errorf("value is so low : %d", rst)
	}

}