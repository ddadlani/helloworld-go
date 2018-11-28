package main

import (
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	os.Setenv("TARGET", "blah")
	SetTarget("test")
	target := os.Getenv("TARGET")
	if target != "test" {
		t.Fail()
	}
}
