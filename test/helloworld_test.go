package test

import (
	"os"
	"testing"

	trg "github.com/ddadlani/helloworld-go/target"
)

func TestHandler(t *testing.T) {
	os.Setenv("TARGET", "blah")
	trg.SetTarget("test")
	target := os.Getenv("TARGET")
	if target != "test" {
		t.Fail()
	}
}
