package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {
	//테스트케이스에 대한 슬라이스 생성
	tests := []testConfig{
		{
			//-h가 Argument로 주어질 경우
			args: []string{"-h"},
			err:  nil,
			//numTimes는 0, printUsage(사용법출력여부)는 true
			config: config{printUsage: true, numTimes: 0},
		},
		{
			//10이 Argument로 주어질 경우
			args: []string{"10"},
			err:  nil,
			//numTimes=10, 사용법출력 false
			config: config{printUsage: false, numTimes: 10},
		},
		{
			//abc가 Argument로 주어질 경우
			args: []string{"abc"},
			//특정 에러 발생
			err: errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			//config값의 경우 false, 0
			config: config{printUsage: false, numTimes: 0},
		},
		{
			//Argument가 1개이상일 경우
			args: []string{"1", "foo"},
			//특정 에러 발생
			err: errors.New("Invalid number of arguments"),
			//false, 0
			config: config{printUsage: false, numTimes: 0},
		},
	}

	for _, tc := range tests {
		//parseArgs메서드 실행
		c, err := parseArgs(tc.args)
		//tc.err가 nil이 아니고 두 에러가 일치하지 않을 때
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		//tc.err가 nil이고 err는 nil이 아닐 때
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got %v\n", err)
		}
		//config의 printUsage가 테스트케이스와 일치하지 않을 때
		if c.printUsage != tc.config.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.config.printUsage, c.printUsage)
		}
		//config의 numTimes가 테스트케이스와 일치하지 않을 때
		if c.numTimes != tc.config.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.config.numTimes, c.numTimes)
		}
	}
}
