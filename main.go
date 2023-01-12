package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var flagNum int

func init() {
	tagmsg := "param"

	flag.IntVar(&flagNum, "n", 0, "例子编号")
	flag.Parse()

	log.Println(tagmsg, "flagNum", flagNum)
}

func main() {
	switch flagNum {
	case 1:
	default:
		log.Println("hello world !")
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sigval := <-sigchan
	log.Println("====>停止程序", sigval)
}
