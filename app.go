package goapp

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fjw201198/simplelog"
)

// App is a common App interface
type App interface {
	Name() string
	Version() string
	Start() bool
}

// Stop will send a SIGQUIT signal to app
func Stop() {
	syscall.SIGQUIT.Signal()
}

// WaitExit a convinience function for wait exit
func WaitExit() {
	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL,
		syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGSEGV)
	for {
		select {
		case s := <-ch:
			processSignal(s)
			break
		}
	}
}

func CreatePidFile(a App) bool {
	var fpath = "__" + a.Name() + "_" + a.Version() + ".pid"
	curPid := os.Getpid()
	fd, err := os.OpenFile(fpath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create pid file failed: %s\n", err)
		return false
	}
	fmt.Fprintf(fd, "%d", curPid)
	fd.Close()
	return true
}

func processSignal(s os.Signal) {
	switch s {
	case syscall.SIGHUP:
		// ignore syscall.SIGHUP
		break
	default:
		simplelog.ExitHook()
		os.Exit(0)
		break
	}
}
