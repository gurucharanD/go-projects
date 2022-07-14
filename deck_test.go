package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != 16{
		t.Errorf("test failed",len(d))
	}

}