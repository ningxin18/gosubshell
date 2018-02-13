package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func argsFunc(s string) []*exec.Cmd {
	var cmdSlice []*exec.Cmd
	n := strings.Index(s, "|")
	if n == -1 {
		args := strings.Fields(s)
		cmdSlice = append(cmdSlice, exec.Command(args[0], args[1:]...))
		return cmdSlice

	}
	args := strings.Split(s, "|")
	for _, v := range args {
		cmd := strings.Fields(v)
		cmdSlice = append(cmdSlice, exec.Command(cmd[0], cmd[1:]...))

	}
	return cmdSlice

}

func execFunc(cmdSlice []*exec.Cmd) error {
	var err error
	max := len(cmdSlice) - 1
	cmdSlice[max].Stdout = os.Stdout
	cmdSlice[max].Stderr = os.Stderr

	for i, cmd := range cmdSlice[:max] {
		if i == max {
			break

		}
		cmdSlice[i+1].Stdin, err = cmd.StdoutPipe()
		if err != nil {
			return err

		}

	}

	for _, cmd := range cmdSlice {
		err := cmd.Start()
		if err != nil {
			return err

		}

	}

	for _, cmd := range cmdSlice {
		err := cmd.Wait()
		if err != nil {
			return err

		}

	}
	return nil

}

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[ningxin@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break

		}

		line := r.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue

		}

		cmdSlice := argsFunc(line)
		err := execFunc(cmdSlice)
		if err != nil {
			fmt.Println(err)

		}

	}

}
