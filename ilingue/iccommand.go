package ilingue

// ICCommand Provides an interface for implementation and standardization of commands
type ICCommand interface {
	ExecuteCommand(param ...string)
}
