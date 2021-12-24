package command

import (
	"os"
	"os/exec"
)

// CClearCmd Provides Windows operating system cmd cleanup command implementation.
type CClearCmdw struct {
}

//
func NewCClearCmdw() *CClearCmdw {
	return &CClearCmdw{}
}

// ExecuteCommand Implements the way windows interprets the command request, for cleaning the cmd.
func (ref *CClearCmdw) ExecuteCommand(param ...string) {
	refcmd := exec.Command("cmd", "/c", "cls")
	refcmd.Stdout = os.Stdout
	refcmd.Run()
}
