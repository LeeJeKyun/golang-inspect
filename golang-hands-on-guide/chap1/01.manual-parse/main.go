package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 커맨드라인 인수를 저장하는 구조체
// 반복횟수와 도움말 출력 여부 저장
type config struct {
	numTimes   int
	printUsage bool
}

// 도움말 문자열 저장
var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name you entered <integer> number
of times.`, os.Args[0])

// 커맨드라인 인수 파싱 함수
func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	// 커맨드라인 인수가 1개이상일 경우 에러 반환
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	// 커맨드라인 인수가 help일 경우 설정값 저장 후 리턴
	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	// 입력받은 커맨드라인 인수를 int로 타입변환하여 저장
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

// 도움말 여부 확인 후 도움말 출력 또는 이름을 입력받아 반복 출력
func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"

	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}

	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func main() {
	//커맨드라인 인수의 두번째 요소부터 파싱 및 설정값 반환
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	//커맨드라인 인수 검증
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	//커맨드라인 인수에 따라 코드 실행
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
