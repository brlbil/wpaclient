package wpaclient

import "errors"

// ErrUnknownCmd returned when a UNKNOWN_COMMAND message received
var ErrUnknownCmd = errors.New("unknown command")

// ErrCmdFailed returned when a FAIL message received
var ErrCmdFailed = errors.New("command failed")

// InvalidCmdError returned when Invalid COMMAND message received
type InvalidCmdError struct {
	Cmd string
	Err string
}

func (ie *InvalidCmdError) Error() string { return ie.Cmd + ": " + ie.Err }
