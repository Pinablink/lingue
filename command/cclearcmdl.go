package command

import (
	"os"
	"os/exec"
)

// CClearCmdl  Provides the implementation of the Linux operating system cmd cleanup command.
type CClearCmdl struct {
}

//
func NewCClearCmdl() *CClearCmdl {
	return &CClearCmdl{}
}

// ExecuteCommand Implements the way linux interprets the command request, to clean up the cmd.
func (ref *CClearCmdl) ExecuteCommand(param ...string) {
	refcmd := exec.Command("clear")
	refcmd.Stdout = os.Stdout
	refcmd.Run()
}
