package main

import (
	"fmt"
	"os/exec"

	"github.com/pastdev/ci/command"
)

func main() {
	runner := command.Runner{
		BufferStdout: true,
		PrintStdout:  true,
	}

	result, err := runner.Run("git", "status", "porcelain")
	if err != nil {
		if runErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("---- EXIT ERROR ----\n%s\n-- END EXIT ERROR --\n", runErr)
		} else {
			fmt.Printf("---- ERROR ----\n%s\n-- END ERROR --\n", err)
		}
	}
	fmt.Printf("---- STDOUT ----\n%s-- END STDOUT --\n", result.Stdout.String())
}
