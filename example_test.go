package testutils_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/britannic/testutils"
)

var err = fmt.Errorf("We have a problem Houston!")

// True returns true
func True() bool {
	return true
}

// TestAssert fails the test if the condition is false.
func TestAssert(t *testing.T) {
	Assert(t, True(), "We have a problem Houston!", true)
}

// TestEquals fails the test if exp is not equal to act.
func TestEquals(t *testing.T) {
	Equals(t, true, true)
}

// TestNotEquals passes the test if exp is not equal to act.
func TestNotEquals(t *testing.T) {
	NotEquals(t, 10, 0)
}

// TestOK fails the test if an err is not nil.
func TestOK(t *testing.T) {
	OK(t, nil)
}

// TestOK passes the test if an err is not nil.
func TestNotOK(t *testing.T) {
	NotOK(t, err)
}

// TestFailures fails every test.
func TestFailures(t *testing.T) {
	if True() {
		t.SkipNow()
	}

	Assert(t, True() == false, "We have a problem Houston!", true)
	Equals(t, false, true)
	NotEquals(t, 0, 0)
	OK(t, err)
	err = errors.New("")
	NotOK(t, err)
}
