package subcmd

import "os"

var (
	name = ""
)

func init() {
	Setup()
}

// Extract and remove sub-command from `os.Args`.
// It's automatically called through `init`
func Setup() {
	if name != "" {
		return
	}

	if len(os.Args) > 1 && os.Args[1][0] != '-' {
		name = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
}

// Get sub-command name. Default sub-command name is an empty string.
// Return default if os.Args has only one element OR
// `os.Args[1]` is a flag or option(string starting with a "-")
func Name() string {
	return name
}

// Restore the sub-command to `os.Args` and reset sub-command to empty string.
func Restore() {
	if name == "" {
		return
	}

	os.Args = append(os.Args[:1], append([]string{name}, os.Args[1:]...)...)
	name = ""
}
