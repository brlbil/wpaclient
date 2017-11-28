package wpaclient

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// validate can evolve to more complex function,
// since we are covering very limited return values

func validate(cmd string, buf []byte) error {
	sb := strings.TrimSuffix(string(buf), "\n")

	switch sb {
	case "UNKNOWN COMMAND":
		return ErrUnknownCmd
	case "FAIL":
		return ErrCmdFailed
	}

	ic := fmt.Sprintf("Invalid %s command", cmd)
	if strings.HasPrefix(sb, ic) {
		errs := strings.Replace(sb, ic, "", -1)
		for i := 0; i < len(errs); i++ {
			r := rune(errs[i])
			if 96 < r && r < 123 {
				errs = errs[i:]
				break
			}
		}
		errs = strings.Replace(errs, "\n", " ", -1)
		return &InvalidCmdError{Cmd: cmd, Err: errs}
	}

	// usage message returned
	if strings.HasPrefix(sb, strings.ToLower(cmd)) {
		return &InvalidCmdError{Cmd: cmd}
	}

	if cmd == CmdPing && sb != "PONG" {
		return errors.Wrapf(ErrCmdFailed, "expected PONG got %s", sb)
	}

	return nil
}
