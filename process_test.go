package main

import (
	"testing"

	"github.com/spf13/afero"
)

var (
	fs afero.Fs
)

func init() {
	// init in-memory file system
	fs = afero.NewMemMapFs()

	// populate in-memory file system
	fs.MkdirAll("/proc/1", 0755)
	fs.MkdirAll("/proc/fs", 0755)
	afero.WriteFile(fs, "/proc/1/comm", []byte("bash\n"), 0644)
	afero.WriteFile(fs, "/proc/1/cmdline", []byte("bash\x00/foo\x00"), 0644)

}

func TestNewProcess(t *testing.T) {
	stat, _ := fs.Stat("/proc/1")

	// init the process
	process := NewProcess(fs, stat)

	// pid should be populated
	if process.Pid != 1 {
		t.Fatalf("expected pid to be 1, got: %d", process.Pid)
	}
}

func TestGetPidFromDir(t *testing.T) {
	stat, _ := fs.Stat("/proc/1")
	if pid := GetPidFromDir(stat); pid != 1 {
		t.Fatalf("expected 1, got: %d", pid)
	}

	stat, _ = fs.Stat("/proc/fs")
	if pid := GetPidFromDir(stat); pid != 0 {
		t.Fatalf("expected 0, got: %d", pid)
	}
}

func TestGetCommFromDir(t *testing.T) {
	stat, _ := fs.Stat("/proc/1")

	if comm := GetCommFromDir(fs, stat); comm != "bash" {
		t.Fatalf("expected \"bash\", got: %q", comm)
	}
}

func TestGetCmdlineFromDir(t *testing.T) {
	stat, _ := fs.Stat("/proc/1")

	if cmdline := GetCmdlineFromDir(fs, stat); cmdline != "bash /foo" {
		t.Fatalf("expected \"bash /foo\", got: %q", cmdline)
	}
}
