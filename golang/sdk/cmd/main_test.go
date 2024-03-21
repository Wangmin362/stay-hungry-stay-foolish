package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"testing"
)

// 一直执行，知道命令结束才打印命令的输出结果
func TestExecCmd(t *testing.T) {
	cmd := exec.Command("ping", "www.baidu.com", "-t")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", string(output))
}

// 测试执行命令，并且实时输出命令的结果，而不是等到命令执行结束之后返回结果
func TestRealTimeCmd(t *testing.T) {
	cmd := exec.Command("ping", "www.baidu.com", "-t")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
