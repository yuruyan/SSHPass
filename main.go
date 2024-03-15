package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("\tsshpass.exe username@ip password")
		os.Exit(-1)
		return
	}
}

func main() {
	cmd := exec.Command("ssh", os.Args[1])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	time.Sleep(time.Second)
	robotgo.TypeStr(os.Args[2] + "\n")
	cmd.Wait()
}
