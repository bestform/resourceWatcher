package main

import (
	"bytes"
	"testing"
)

func init() {
	initDb(new(bytes.Buffer))
}

func TestSetContentForResource(t *testing.T) {
	if hasContentForResource("foo") {
		t.Errorf("There should be no content until set")
	}

	setContentForResource("foo", "bar")
	if !hasContentForResource("foo") {
		t.Errorf("There should be content after setting it")
	}
}

func TestHasContentChanged(t *testing.T) {
	setContentForResource("foo", "bar")
	if hasContentChanged("foo", "bar") {
		t.Errorf("There should be no change with same content")
	}
	if !hasContentChanged("foo", "bar2") {
		t.Errorf("There should be a change with different content")
	}
}
