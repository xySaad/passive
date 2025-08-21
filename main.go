package main

import (
	"context"
	"fmt"
	"os"
	"passive/apis"
	fullname "passive/apis/fullname"
	"passive/apis/ip"
	"passive/apis/username"

	"passive/help"

	"github.com/chromedp/chromedp"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(help.USAGE)
		return
	}

	flag := args[0]
	if flag == "--help" {
		fmt.Println(help.HELP_FLAG)
		return
	}
	usage, exec := HandleFlag(flag)
	if exec == nil {
		fmt.Printf(usage, flag)
		return
	}

	if len(args) == 1 {
		fmt.Println(usage)
		return
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	parent, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	exec(parent, args[1])
}

func HandleFlag(flag string) (usage string, fn apis.Executor) {
	switch flag {
	case "-fn":
		return fullname.MISSING_INPUT, fullname.Exec
	case "-ip":
		return ip.MISSING_INPUT, ip.Exec
	case "-u":
		return username.MISSING_INPUT, username.Exec
	}
	return help.INVALID_FLAG, nil
}
