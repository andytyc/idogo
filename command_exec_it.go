package main

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CmdExecIT(ctx context.Context) {
	tagmsg := "执行交互命令"

	pwdPath, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	cmdDir := filepath.Join(pwdPath, ".")
	log.Println(tagmsg, "当前目录", pwdPath)

	funcDo := func(input string) (cmdName string, args []string, ok bool) {
		// log.Println(tagmsg, "输入", input)

		inputArgs := strings.Split(input, " ")
		if len(inputArgs) <= 0 {
			log.Println(tagmsg, "执行命令为空,重新输入", input)
			return
		}
		cmdName = inputArgs[0]
		if len(inputArgs) > 1 {
			args = inputArgs[1:]
		} else {
			args = inputArgs[:0]
		}
		ok = true

		// log.Println(tagmsg, "输入成功", "命令", cmdName, "参数", args)
		return
	}

	log.Println(tagmsg, "输入格式:\n命令 参数1 参数2,...")

	inputReader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Println(tagmsg, "请输入:")

			input, err := inputReader.ReadString('\n')
			if err != nil {
				log.Println(tagmsg, "输入异常,退出", err)
				os.Exit(0)
			}
			if len(input) <= 1 {
				// log.Println(tagmsg, "仅回车操作,重新输入") // 仅回车则忽略
				continue
			}
			input = strings.TrimRight(input, "\n")

			cmdName, args, ok := funcDo(input)
			if !ok {
				continue
			}

			if cmdName == "exit" || cmdName == "q" {
				log.Println(tagmsg, "退出")
				os.Exit(0)
			} else if cmdName == "cd" {
				if len(args) > 0 {
					cmdDirCd := args[0]
					cmdDirNew := filepath.Join(cmdDir, cmdDirCd)

					finfo, err := os.Stat(cmdDirNew)
					if err != nil && !os.IsExist(err) {
						log.Println(tagmsg, "路径不存在,重新输入", input)
						continue
					}
					if !finfo.IsDir() {
						log.Println(tagmsg, "路径非目录,重新输入", input)
						continue
					}

					cmdDir = cmdDirNew
					log.Println(tagmsg, "执行成功", cmdDir)
					continue
				} else {
					log.Println(tagmsg, "切换目录路径为空,重新输入", input)
					continue
				}
			}
			// log.Println(tagmsg, "执行命令", cmdDir, cmdName, args)

			cmd := exec.Command(cmdName, args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = os.Stderr
			cmd.Dir = cmdDir
			err = cmd.Start()
			if err != nil {
				log.Println(tagmsg, "执行失败", err)
				continue
			}
			err = cmd.Wait()
			if err != nil {
				log.Println(tagmsg, "执行失败", err)
				continue
			}
			result := strings.TrimRight(out.String(), "\n")
			log.Println(tagmsg, "执行成功\n"+result)
		}
	}
}

/*
func CmdExecIT(ctx context.Context) {
	tagmsg := "执行交互命令"

	funcDo := func(input string) (cmdDir, cmdName string, args []string, ok bool) {
		// log.Println(tagmsg, "输入", input)

		input = strings.TrimRight(input, "\n")
		inputArgs := strings.Split(input, " ")
		if len(inputArgs) <= 0 {
			log.Println(tagmsg, "执行目录为空,重新输入", input)
			return
		}
		for i, arg := range inputArgs {
			if i == 0 {
				cmdDir = arg
			} else if i == 1 {
				cmdName = arg
			}
		}
		if len(inputArgs) > 2 {
			args = inputArgs[2:]
		} else {
			args = inputArgs[:0]
		}
		ok = true

		// log.Println(tagmsg, "输入成功", "目录", cmdDir, "命令", cmdName, "参数", args)
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Println(tagmsg, "请输入,格式: 目录,命令,参数1,参数2,...")
			input, err := inputReader.ReadString('\n')
			if err != nil {
				log.Println(tagmsg, "输入异常,重新输入", err)
				continue
			}
			if len(input) <= 1 {
				// log.Println(tagmsg, "仅回车操作,重新输入") // 仅回车则忽略
				continue
			}

			cmdDir, cmdName, args, ok := funcDo(input)
			if !ok {
				continue
			}
			if cmdName == "" {
				log.Println(tagmsg, "输入命令为空,重新输入")
				continue
			} else if cmdName == "exit" || cmdName == "q" {
				log.Println(tagmsg, "退出")
				os.Exit(0)
			}

			cmd := exec.Command(cmdName, args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = os.Stderr
			cmd.Dir = cmdDir
			err = cmd.Start()
			if err != nil {
				log.Println(tagmsg, "执行失败", err)
				continue
			}
			err = cmd.Wait()
			if err != nil {
				log.Println(tagmsg, "执行失败", err)
				continue
			}
			log.Println(tagmsg, "执行成功\n", out.String())
		}
	}
}
*/
