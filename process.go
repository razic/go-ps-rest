package main

import (
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
func NewProcess(fs afero.Fs, dir os.FileInfo) *Process {
	return &Process{
		Pid:  GetPidFromDir(dir),
		Comm: GetCommFromDir(fs, dir),
	}
}

// GetPidFromDir gets pid from a dir
func GetPidFromDir(dir os.FileInfo) int {
	pid, _ := strconv.Atoi(filepath.Base(dir.Name()))

	return pid
}

// GetCommFromDir gets the comm from a dir
func GetCommFromDir(fs afero.Fs, dir os.FileInfo) string {
	bytes, err := afero.ReadFile(fs, "/proc/"+dir.Name()+"/comm")

	if err != nil {
		return ""
	}

	return string(bytes[:len(bytes)-1])
}
