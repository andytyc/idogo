package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var flagNum int

func init() {
	tagmsg := "param"

	flagNumMap := map[int]string{
		0: "hello world",
		1: "执行命令",
		2: "执行交互命令",
	}
	flagNumMsg := func() string {
		msg := "\n"
		for key, val := range flagNumMap {
			msg += fmt.Sprintf("%d  %s\n", key, val)
		}
		msg = strings.TrimRight(msg, "\n")
		return msg
	}
	flag.IntVar(&flagNum, "n", 0, "案例编号，选项:"+flagNumMsg())

	flag.Parse()
	log.Println(tagmsg, "flag", flag.NFlag())
	flag.VisitAll(func(arg *flag.Flag) {
		if arg.Name == "n" {
			log.Println(arg.Name, arg.Value, "案例编号")
		} else {
			log.Println(arg.Name, arg.Value, arg.Usage)
		}
	})
	log.Println(tagmsg, "args", flag.NArg(), flag.Args())

	log.Println(tagmsg, "加载完毕\n****\n启动\n****")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch flagNum {
	case 1:
		CmdExec()
	case 2:
		CmdExecIT(ctx)
	default:
		log.Println("hello world !")
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sigval := <-sigchan
	log.Println("====>停止程序", sigval)
}
