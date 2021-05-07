package util

import "testing"

func TestGenMobile(t *testing.T) {
	mobile := GenMobile()
	if 11 != len(mobile) {
		t.Errorf("got %d want 11 given, %v", len(mobile), mobile)
	} else {
		t.Logf("got %d want 11 given, %v", len(mobile), mobile)
	}
}
