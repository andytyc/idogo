package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
)

func CmdExec() {
	tagmsg := "执行命令"

	if flag.NArg() < 2 {
		log.Println(tagmsg, "缺少操作目录和命令，参考:\n操作目录 命令 参数1 参数2 ...")
		os.Exit(0)
	}
	args := flag.Args()
	cmdDir := args[0]
	cmdName := args[1]
	args = args[2:]

	cmd := exec.Command(cmdName, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Dir = cmdDir
	err := cmd.Start()
	if err != nil {
		log.Println(tagmsg, "执行失败", err)
		os.Exit(0)
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(tagmsg, "执行失败", err)
		os.Exit(0)
	}
	log.Println(tagmsg, "执行成功\n", out.String())
}
