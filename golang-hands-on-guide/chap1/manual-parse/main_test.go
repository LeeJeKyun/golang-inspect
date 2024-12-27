package main

import (
	"os/exec"
	"testing"
)

func TestMainFunction(t *testing.T) {

	//TODO 테스트케이스 추가 및 케이스에 따른 error발생, output변경 추가
	tests := []struct {
		command string
		args    []string
		input   string
		output  string
		err     error
	}{
		{
			command: "./application",
			args:    []string{"3"},
			input:   "LeeJeKyun",
			output: `Your name please? Press the Enter key when done.
Nice to meet you LeeJeKyun
Nice to meet you LeeJeKyun
Nice to meet you LeeJeKyun
`,
		},
		//		{
		//			command: "./application",
		//			output: `Invalid number of arguments
		//		Usage: ./application <integer> [-h|--help]
		//
		//		A greeter application which prints the name you entered <integer> number of times
		//`,
		//		},
		//{
		//	command: "./application",
		//	args:    []string{"0"},
		//	err:     errors.New("Most specify a number greater than 0"),
		//},
	}

	for _, test := range tests {
		command := exec.Command(test.command, test.args...)
		pipe, _ := command.StdinPipe()
		pipe.Write([]byte(test.input))
		pipe.Close()
		output, _ := command.CombinedOutput()
		if string(output) != test.output {
			t.Fatalf("Error Occured!!!")
		}
	}

}
