package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// 커맨드라인 인수를 저장하는 구조체
// 반복횟수와 도움말 출력 여부 저장
type config struct {
	numTimes int
	name     string
}

// 에러 전역변수 생성
var errPosArgsSpecified = errors.New("Positional arguments specified")
var errInvalidPosArgSpecified = errors.New("More than one positional argument specified")

// 커맨드라인 인수 파싱 함수
func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}

	// 애플리케이션 사용법에서 보여주는 애플리케이션 명 및 파싱 도중 에러처리 방식 설정
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)

	// FlagSet 객체의 진단 메시지 혹은 출력 메시지를 작성하는데 사용되는 Writer 지정
	fs.SetOutput(w)

	// FlagSet 사용법 메시지 정의
	fs.Usage = func() {
		var usageString = `
A greeter application which prints the name you entered a specified number of times.

Usage of: %s `
		fmt.Fprintf(w, usageString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	// IntVar(저장할 변수의 주소값, 옵션의 이름, 기본값, 옵션의 목적을 나타내는 문자열)
	// usage 매개변수는 사용법 출력 시 자동으로 보이게 됨.
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	// 위치인수 개수가 옳지 않을 경우
	// 에러 반환
	if fs.NArg() > 1 {
		return c, errPosArgsSpecified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}

	return c, nil
}

// 도움말 여부 확인 후 도움말 출력 또는 이름을 입력받아 반복 출력
func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}
	greetUser(c, w)
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

func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func main() {
	//커맨드라인 인수의 두번째 요소부터 파싱 및 설정값 반환
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		// 위치인수 관련 에러가 발생했을 경우에만 에러를 출력
		if errors.Is(err, errPosArgsSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}

	//커맨드라인 인수 검증
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	//커맨드라인 인수에 따라 코드 실행
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
