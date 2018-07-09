package command

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

type Runner struct {
	BufferStdout bool
	BufferStderr bool
	PipeStdout   io.Writer
	PipeStderr   io.Writer
	PrintStdout  bool
	PrintStderr  bool
}

type RunResult struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
}

func Run(command string, args ...string) error {
	_, err := Runner{}.Run(command, args...)
	return err
}

func (r Runner) Run(command string, args ...string) (RunResult, error) {
	return run(r, command, args...)
}

func run(r Runner, command string, args ...string) (RunResult, error) {
	cmd := exec.Command(command, args...)

	result := RunResult{}

	if r.BufferStdout || r.PipeStdout != nil || r.PrintStdout {
		writers := []io.Writer{}

		if r.PipeStdout != nil {
			writers = append(writers, r.PipeStdout)
		}

		if r.PrintStdout {
			writers = append(writers, os.Stdout)
		}

		if r.BufferStdout {
			writers = append(writers, &result.Stdout)
		}

		cmd.Stdout = io.MultiWriter(writers...)
	}

	if r.BufferStderr || r.PipeStderr != nil || r.PrintStderr {
		writers := []io.Writer{}

		if r.PipeStderr != nil {
			writers = append(writers, r.PipeStderr)
		}

		if r.PrintStderr {
			writers = append(writers, os.Stderr)
		}

		if r.BufferStderr {
			writers = append(writers, &result.Stderr)
		}

		cmd.Stderr = io.MultiWriter(writers...)
	}

	return result, cmd.Run()
}
