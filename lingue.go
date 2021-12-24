package lingue

import (
	"github.com/Pinablink/lingue/command"
	"github.com/Pinablink/lingue/ilingue"
	"github.com/Pinablink/lingue/oslabel"
)

// It provides resources that standardize some Windows and Linux commands.
type Lingue struct {
	mapResource map[oslabel.OSLabel]map[oslabel.OSCommand]ilingue.ICCommand
}

// NewLingue Provides an instance of the Lingue structure
func NewLingue() *Lingue {
	return &Lingue{mapResource: mkResource()}
}

// ExecCommand Execute the requested command according to the operating system provider.
func (lingue *Lingue) ExecCommand(refCommand oslabel.OSCommand) error {

	key, err := oslabel.QOS()

	if err != nil {
		return err
	}

	(lingue.mapResource[key])[refCommand].ExecuteCommand()

	return nil
}

// mkResource Build the map containing the platform commands. This version contains only Linux and Windows.
func mkResource() map[oslabel.OSLabel]map[oslabel.OSCommand]ilingue.ICCommand {
	mapRef := make(map[oslabel.OSLabel]map[oslabel.OSCommand]ilingue.ICCommand)

	var cclearLinux ilingue.ICCommand = command.NewCClearCmdl()
	var cclearWindows ilingue.ICCommand = command.NewCClearCmdw()

	mapCmmLinux := make(map[oslabel.OSCommand]ilingue.ICCommand)
	mapCmmWindows := make(map[oslabel.OSCommand]ilingue.ICCommand)

	mapCmmLinux[oslabel.CLEAR_CMD] = cclearLinux
	mapCmmWindows[oslabel.CLEAR_CMD] = cclearWindows

	mapRef[oslabel.LINUX] = mapCmmLinux
	mapRef[oslabel.WINDOWS] = mapCmmWindows

	return mapRef
}
