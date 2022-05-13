package main

import "testing"

func TestBackup(t *testing.T) {
	backup("/Users/.ctrlc/", "/Users/.ctrlc/ctrlc.db")

}
