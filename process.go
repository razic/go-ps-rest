package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/afero"
)

// Process is a data structure representing a process
type Process struct {
	Pid     int      `json:"pid"`
	Comm    string   `json:"name"`
	Cmdline string   `json:"cmdline"`
	Environ []string `json:"environment,"`
}

// NewProcess takes a filesystem and directory, and populates the struct fields
func NewProcess(fs afero.Fs, dir os.FileInfo) *Process {
	return &Process{
		Pid:     GetPidFromDir(dir),
		Comm:    GetCommFromDir(fs, dir),
		Cmdline: GetCmdlineFromDir(fs, dir),
		Environ: GetEnvironFromDir(fs, dir),
	}
}

// GetPidFromDir gets pid from a dir
func GetPidFromDir(dir os.FileInfo) int {
	pid, _ := strconv.Atoi(filepath.Base(dir.Name()))

	return pid
}

// GetCommFromDir gets the comm from a dir
func GetCommFromDir(fs afero.Fs, dir os.FileInfo) string {
	byteArr, err := afero.ReadFile(fs, "/proc/"+dir.Name()+"/comm")

	if err != nil {
		return ""
	}

	return string(byteArr[:len(byteArr)-1])
}

// GetCmdlineFromDir gets the cmdline from a dir
func GetCmdlineFromDir(fs afero.Fs, dir os.FileInfo) string {
	byteArr, err := afero.ReadFile(fs, "/proc/"+dir.Name()+"/cmdline")

	if err != nil {
		return ""
	}

	return string(bytes.Replace(byteArr[:len(byteArr)-1], []byte("\x00"), []byte(" "), -1))
}

// GetEnvironFromDir gets the environment from a dir
func GetEnvironFromDir(fs afero.Fs, dir os.FileInfo) []string {
	byteArr, err := afero.ReadFile(fs, "/proc/"+dir.Name()+"/environ")

	if err != nil {
		return []string{}
	}

	pieces := bytes.Split(byteArr[:len(byteArr)-1], []byte("\x00"))
	slice := make([]string, len(pieces))

	for i, piece := range pieces {
		slice[i] = string(piece)
	}

	return slice
}
