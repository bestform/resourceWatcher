package main

import (
	"bytes"
	"testing"
)

func init() {
	initDb(new(bytes.Buffer))
}

func TestSetContentForResource(t *testing.T) {
	if HasContentForResource("foo") {
		t.Errorf("There should be no content until set")
	}

	SetContentForResource("foo", "bar")
	if !HasContentForResource("foo") {
		t.Errorf("There should be content after setting it")
	}
}

func TestHasContentChanged(t *testing.T) {
	SetContentForResource("foo", "bar")
	if HasContentChanged("foo", "bar") {
		t.Errorf("There should be no change with same content")
	}
	if !HasContentChanged("foo", "bar2") {
		t.Errorf("There should be a change with different content")
	}
}
