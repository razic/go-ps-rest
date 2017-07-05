package main

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/afero"
)

// Process is a data structure representing a process
type Process struct {
	Pid  int
	Comm string
}

// NewProcess takes a filesystem and directory, and populates the struct fields
func NewProcess(fs afero.Fs, dir os.FileInfo) (*Process, error) {
	process := &Process{Pid: GetPidFromDir(dir.Name())}

	if process.Pid == 0 {
		return nil, errors.New("unable to obtain pid")
	}

	return process, nil
}

// GetPidFromDir gets pid from a dirname
func GetPidFromDir(dir string) int {
	pid, err := strconv.Atoi(filepath.Base(dir))

	if err != nil {
		return 0
	}

	return pid
}
