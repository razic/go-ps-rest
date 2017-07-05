package main

import (
	"os"
	"regexp"

	"github.com/spf13/afero"
)

// ProcessLister implements the ProcessLister interface
type ProcessLister interface {
	List() []*Process
}

// ProcessList is an array of pointers to processes
type ProcessList struct {
	fs afero.Fs
}

// List returns process list
func (l *ProcessList) List() []*Process {
	return []*Process{}
}

// GetProcDirs returns a list of proc dirs
func (l *ProcessList) GetProcDirs() ([]os.FileInfo, error) {
	// get list of all directories in /proc
	dirsA, err := afero.ReadDir(l.fs, "/proc")

	if err != nil {
		return nil, err
	}

	// create a new slice, using same backing array
	dirsB := dirsA[:0]

	// filter out the directories not matching a numeric pattern
	for _, dir := range dirsA {
		if ok, _ := regexp.Match(`^\d+$`, []byte(dir.Name())); ok {
			dirsB = append(dirsB, dir)
		}
	}

	return dirsB, nil
}
