package main

import (
	"os"
	"path/filepath"
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
	// get a list of all proc dirs
	procDirs, err := l.GetProcDirs()

	// initialize an array to hold the proccess objects
	plist := []*Process{}

	// error checking
	if err != nil {
		return plist
	}

	// loop over the proc dirs to initialize a process object for each
	for _, dir := range procDirs {
		if process := NewProcess(l.fs, dir); process != (&Process{}) && filepath.Base(os.Args[0]) != process.Comm {
			plist = append(plist, process)
		}
	}

	return plist
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
