package std

import "testing"

func TestDemoRand(t *testing.T) {
	if !DemoRand() {
		t.Error("failed...")
	}
}
