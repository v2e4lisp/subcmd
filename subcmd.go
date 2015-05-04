package subcmd

import "os"

var (
	Name = ""
)

func init() {
	if len(os.Args) > 1 && os.Args[1][0] != '-' {
		Name = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
}
