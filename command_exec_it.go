package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
)

func CmdExecIT(ctx context.Context) {
	tagmsg := "执行交互命令"

	inputReader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			input, err := inputReader.ReadString('\n')
			if err != nil {
				log.Println(tagmsg, "输入异常", err)
				continue
			}
			if len(input) <= 1 {
				log.Println(tagmsg, "输入为空，请继续输入", err)
				continue
			}
			log.Println(tagmsg, "输入", input)
			fmt.Sscan()
		}
	}
}
