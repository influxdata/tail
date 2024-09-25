// +build linux darwin freebsd netbsd openbsd aix

package tail

import (
	"os"
)

func OpenFile(name string) (file *os.File, err error) {
	return os.Open(name)
}
