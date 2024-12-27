package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

// 사용법에 대한 문자열
var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name you entered <integer> number of times`, os.Args[0])

// 사용법 출력
func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

// 명령줄인자 검증(Arguments가 0보다 작거나 같을 경우 에러 반환)
func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Most specify a number greater than 0")
	}
	return nil
}

// Arguments 파싱
func parseArgs(args []string) (config, error) {

	//Arguments로 입력된 numTimes 파싱
	var numTimes int
	var err error
	c := config{}

	//Arguments의 갯수가 1개가 아닌 경우 모두 에러 발생
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	//Argument가 -h나 --help일 경우 printUsage = true 설정(방법 설명 설정)
	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	// 위의 에러가 모두 발생하지 않을 경우 Arguments의 string값을 int로 변환
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	// config 객체의 numTimes를 설정
	c.numTimes = numTimes

	return c, nil
}

// 이름을 입력받는 메서드, 반환값은 이름, 에러
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"

	//매개변수로 입력된 io.Writer 인터페이스 구현체에 msg 값을 출력
	fmt.Fprintf(w, msg)
	//매개변수로 입력된 io.Reader 인터페이스 구현체를 이용한 Scanner 생성
	scanner := bufio.NewScanner(r)
	//한줄 스캔
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	//위에서 스캔해둔 한줄짜리 바이트데이터를 string으로 변환 후 반환
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}

// 이름을 출력하는 메서드
func greetUser(c config, name string, w io.Writer) {
	//Format에 맞게 string을 생성
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		//매개변수로 입력된 io.Writer 인터페이스 구현체에 msg 값을 출력
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {

	// printUsage가 true일 경우 사용법을 출력
	if c.printUsage {
		printUsage(w)
		return nil
	}

	// 이름을 반환받은 후
	name, err := getName(r, w)
	if err != nil {
		return err
	}

	//이름에 대해 횟수만큼 반복
	greetUser(c, name, w)
	return nil
}

func main() {

	// numTimes값을 추출한 config(c) 반환
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	// 변환한 numTimes값을 검증(0보다 커야함.)
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	// 검증된 값을 이용하여 runCmd 실행
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
