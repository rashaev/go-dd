package main

import (
	"testing"
)

func TestGoddWithoutOffsetLimit(t *testing.T) {
	if Godd("testdata/testfile", "testdata/testfiledst", 0, 0) != 913 {
		t.Errorf("source file does't match destination file after copy operation")
	}
}

func TestGoddWithOffset(t *testing.T) {
	if Godd("testdata/testfile", "testdata/testfiledst_offset", 3, 0) != 910 {
		t.Errorf("offset flag does't work")
	}
}

func TestGoddWithLimit(t *testing.T) {
	if Godd("testdata/testfile", "testdata/testfiledst_limit", 3, 2) != 2 {
		t.Errorf("limit flag does't work")
	}
}