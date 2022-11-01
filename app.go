package main

import (
	"bytes"
	"context"
	"fmt"
	. "gtools/models"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetProcessListByPortNumber(portNumber string) []CmdStruct {
	fmt.Println("GetProcessListByPortNumber")

	cmd := exec.Command("lsof", "-i", "tcp:"+portNumber)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalln("err on start")
	}

	err = cmd.Wait()
	cmdRes := out.String()

	s := strings.Split(cmdRes, "\n")
	var cmdResList []CmdStruct
	for k, v := range s {

		if k == 0 || v == "" {
			continue
		}

		for _, v1 := range strings.Split(v, "\n") {
			v2 := strings.Fields(v1)

			cmdStruct := CmdStruct{Command: v2[0], Pid: v2[1], User: v2[2], IpProtocol: v2[4], Connection: v2[8]}
			cmdResList = append(cmdResList, cmdStruct)
		}

	}
	return cmdResList
}

func (a *App) KillProcessByPortNumber(pidStr string) bool {
	fmt.Println("KillProcessByPortNumber")
	pid, _ := strconv.Atoi(pidStr)
	_ = syscall.Kill(pid, syscall.SIGKILL)
	return true
}

func BytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
