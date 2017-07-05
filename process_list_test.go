package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestProccessListGetProcDirs(t *testing.T) {
	// initialize an in-memory file system
	fs := afero.NewMemMapFs()

	// initialize the process list with the in memory file system
	processList := &ProcessList{fs: fs}

	// populate the in-memory file system with /proc
	fs.MkdirAll("/sock/1", 0755)
	fs.MkdirAll("/proc/1", 0755)
	fs.MkdirAll("/proc/2", 0755)
	fs.MkdirAll("/proc/3", 0755)
	fs.MkdirAll("/proc/a", 0755)
	fs.MkdirAll("/proc/b", 0755)
	fs.MkdirAll("/proc/c", 0755)

	dirs, err := processList.GetProcDirs()
	occs := make(map[string]int)

	if err != nil {
		t.Fatal("unexpected error")
	}

	for _, dir := range dirs {
		occs[dir.Name()]++
	}

	if len(occs) != 3 {
		t.Fatalf("expected only 3 dirs, got %d", len(occs))
	}
}
