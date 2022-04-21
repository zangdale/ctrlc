package utils

import "testing"

func TestHomeDir(t *testing.T) {
	s, err := HomeDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
	s, err = HomeDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
