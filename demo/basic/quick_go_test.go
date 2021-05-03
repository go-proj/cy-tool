package basic

import (
	"fmt"
	"testing"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestDemoIf(t *testing.T) {
	DemoIf()
}

func TestDemoStruct(t *testing.T) {
	DemoStruct()
}

func TestDemoInterface(t *testing.T) {
	DemoInterface()
}

func TestDemoErr(t *testing.T) {
	DemoError()
}