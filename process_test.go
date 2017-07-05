package main

import (
	"testing"

	"github.com/spf13/afero"
)

var fs afero.Fs

func init() {
	// init in-memory file system
	fs = afero.NewMemMapFs()

	// populate in-memory file system
	fs.MkdirAll("/proc/1", 0755)

}

func TestNewProcess(t *testing.T) {
	// stat dir
	stat, err := fs.Stat("/proc/1")

	// check for errors
	if err != nil {
		t.Fatal(err)
	}

	// init the process
	process, err := NewProcess(fs, stat)

	// check for errors
	if err != nil {
		t.Fatal(err)
	}

	// pid should be populated
	if process.Pid != 1 {
		t.Fatalf("expected pid to be 1, got: %d", process.Pid)
	}
}

func TestGetPidFromDir(t *testing.T) {
	if pid := GetPidFromDir("/proc/1024"); pid != 1024 {
		t.Fatalf("expected 1024, got: %d", pid)
	}

	if pid := GetPidFromDir("/proc/abc"); pid != 0 {
		t.Fatalf("expected 0, got: %d", pid)
	}
}
