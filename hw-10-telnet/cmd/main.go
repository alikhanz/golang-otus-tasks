package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gookit/color"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "10s", "10")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)

	checkEmptyArg(host, "host")
	checkEmptyArg(port, "port")

	addr := strings.Builder{}
	addr.Write([]byte(host))
	addr.Write([]byte(":"))
	addr.Write([]byte(port))

	timeoutDuration, err := time.ParseDuration(timeout)
	exitCheckErr(err)

	conn, err := net.DialTimeout("tcp", addr.String(), timeoutDuration)
	exitCheckErr(err)
	writeSuccessText(fmt.Sprintf("Success connect: %s", addr.String()))

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go read(conn, wg)
	go write(conn)

	wg.Wait()
}

func write(conn net.Conn) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		_, err := conn.Write(s.Bytes())
		if err != nil {
			writeError(err)
			break
		}
	}
}

func read(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	s := bufio.NewScanner(conn)

	for s.Scan() {
		_, _ = fmt.Fprintln(os.Stdout, s.Text())
	}
}

func checkEmptyArg(arg, name string) {
	if arg == "" {
		exitText(fmt.Sprintf("%s not presented. Usage: telnet 127.0.0.1 1234", name))
	}
}

func exitCheckErr(err error) {
	if err != nil {
		exitError(err)
	}
}

func writeError(err error) {
	writeErrorText(err.Error())
}

func writeErrorText(msg string) {
	color.SetOutput(os.Stderr)
	color.Error.Println(msg)
}

func writeSuccessText(msg string) {
	color.SetOutput(os.Stdout)
	color.Success.Println(msg)
}

func exitError(err error) {
	exitText(err.Error())
}

func exitText(msg string) {
	writeErrorText(msg)
	os.Exit(1)
}