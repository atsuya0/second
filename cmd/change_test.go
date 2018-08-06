package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChange(t *testing.T) {
	patterns := []struct {
		cmd    string
		output string
		err    string
	}{
		{cmd: "second change bin", output: "/usr/bin", err: "%s is not %s"},
		{cmd: "second change xd", output: "/etc/X11/xorg.conf.d", err: "%s is not %s"},
	}

	path, err := filepath.Abs("./testdata/test.json")
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("SECOND_CMD_PATH", path); err != nil {
		log.Fatalln(err)
	}

	var buffer *bytes.Buffer
	cmd := rootCmd()

	for _, p := range patterns {
		buffer = new(bytes.Buffer)
		cmd.SetOutput(buffer)

		args := strings.Split(p.cmd, " ")
		fmt.Printf("args: %+v\n", args)
		cmd.SetArgs(args[1:])
		cmd.Execute()

		result := buffer.String()
		if p.output != result {
			t.Errorf(p.err, result, p.output)
		}
	}
}
