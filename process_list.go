package main

import "github.com/spf13/afero"

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
