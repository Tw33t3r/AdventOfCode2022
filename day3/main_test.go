package main

import (
	"testing"
)

func TestEncode(t *testing.T) {
	result := encodeText('a')
	if result != 1 {
		t.Fatalf("a did not equal 1. Got: %v", result)
	}
	capResult := encodeText('A')
	if capResult != 27 {
		t.Fatalf("A did not equal 27. Got: %v", capResult)
	}
}

func TestSplit(t *testing.T) {
	first, second := splitString("vJrwpWtwJgWrhcsFMMfFFhFp")
	if first != "vJrwpWtwJgWr" || second != "hcsFMMfFFhFp" {
		t.Fatalf("incorrect split")
	}
}
