package main

import (
	"os/exec"
	"syscall"
	"testing"
)

type runTest struct {
	command  string
	args     []string
	input    string
	exitCode int
}

func TestRun(t *testing.T) {

	tests := []runTest{
		// 커맨드라인 인수 갯수 에러
		{
			command:  "go",
			args:     []string{"run", "main.go", "5", "123"},
			input:    "John Doe\n",
			exitCode: 1,
		},
		// 입력값 에러(길이 0)
		{
			command:  "go",
			args:     []string{"run", "main.go", "5"},
			input:    "\n",
			exitCode: 1,
		},
		// 정상 처리
		{
			command:  "go",
			args:     []string{"run", "main.go", "5"},
			input:    "John Doe\n",
			exitCode: 0,
		},
	}

	for _, tc := range tests {
		cmd := exec.Command(tc.command, tc.args...)
		pipe, err := cmd.StdinPipe()
		if err != nil {
			t.Fatalf("Error creating StdinPipe: %v\n", err)
		}
		defer pipe.Close()
		pipe.Write([]byte(tc.input))
		_, err = cmd.Output()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				status := exitError.Sys().(syscall.WaitStatus)
				exitStatus := status.ExitStatus()
				if exitStatus != tc.exitCode {
					t.Fatalf("Expected exit code: %d, Got: %d\n", tc.exitCode, exitStatus)
				}
			}
		}
	}
}
