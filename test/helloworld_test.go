package test

import (
	"os"
	"testing"

	"github.com/ddadlani/helloworld-go/target"
)

func TestHandler(t *testing.T) {
	os.Setenv("TARGET", "blah")
	target.SetTarget("test")
	target = os.Getenv("TARGET")
	if target != "test" {
		t.Fail()
	}
}
