package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

// command is a global variable, so it is not passed to RunCommand
func RunCommand() (string, error) {
	if len(command) == 0 {
		return "", nil
	}
	// TODO handle commands with strings that have spaces in them
	args := strings.Fields(string(command))
	command = nil

	name := ""
	if len(args) > 1 {
		name, args = args[0], args[1:]
	} else {
		name = args[0]
		args = nil
	}
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}

	bts, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return string(bts), nil
}
