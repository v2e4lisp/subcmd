package subcmd

import (
        "os"
        "testing"
)

func TestName(t *testing.T) {
        var name string

        // normal case
        os.Args = []string{"cmd", "subcmd", "-option", "value"}
        Setup()
        name = Name()
        if Name() != "subcmd" {
                t.Error("NORMAL CASE expected subcmd, got", name)
        }
        Restore()

        // only one arg
        os.Args = []string{"cmd"}
        Setup()
        name = Name()
        if Name() != "" {
                t.Error("ONLY ONE ARG expected subcmd, got", name)
        }
        Restore()

        // os.Args[1] is a flag
        os.Args = []string{"cmd", "-a"}
        Setup()
        name = Name()
        if Name() != "" {
                t.Error("os.Args[1] IS A FLAG expected subcmd, got", name)
        }
        Restore()
}
