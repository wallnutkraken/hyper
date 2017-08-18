//+build !windows
package hyper

import (
	"github.com/takama/daemon"
	"os"
)

func IsSupported() bool {
	_, err := os.Stat("/usr/lib/systemd")
	return !os.IsNotExist(err)
}

