// to test, create file names _test.go
// run, go test


package main

import "testing"

func TestNewDeck(t *testing.T){ //name of the fn hsould be T for Go to identify as a test fn
	d:= newdeck()
	if len(d)!=16 {
		t.ErrorF("Expected length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.ErrorF("Frst str of slice is not valid")
	}
}