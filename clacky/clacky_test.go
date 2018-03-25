package clacky

import (
	"testing"
)

func TestClacky(t *testing.T) {
	c := New("f35f52a230350827fc6335388b078c5102cc0776-1")
	longURL, err := c.Short("https://google.com")
	if err != nil {
		t.Errorf("clacky short error: %v", err)
	}
	expected := "clacky.org/JM0e"
	if longURL != expected {
		t.Errorf("got %v\nwant %v", longURL, expected)
	}
}
