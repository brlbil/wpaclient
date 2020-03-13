package wpaclient

type constError string

func (c constError) Error() string { return string(c) }

// ErrUnknownCmd returned when a UNKNOWN_COMMAND message received
var ErrUnknownCmd = constError("unknown command")

// ErrCmdFailed returned when a FAIL message received
var ErrCmdFailed = constError("command failed")

// InvalidCmdError returned when Invalid COMMAND message received
type InvalidCmdError struct {
	Cmd string
	Err string
}

func (ie *InvalidCmdError) Error() string { return ie.Cmd + ": " + ie.Err }
