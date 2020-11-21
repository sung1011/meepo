package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/tidwall/gjson"
)

// ExistDir _
func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

//RunCmd 执行命令
func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var args string
	if len(cmd.Args) > 0 {
		for _, arg := range cmd.Args[1:] {
			args += " " + arg
		}
	}
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Panicf(fmt.Sprintf("cmd.Start() failed with '%s'\n", err))
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, stdoutIn)
	if err != nil {
		log.Panic(err)
	}
	str := buf.String()
	if !gjson.Valid(str) {
		_, errStdout = io.Copy(stdout, bytes.NewReader(buf.Bytes()))
	} else {
		var outBuf bytes.Buffer
		err := json.Indent(&outBuf, buf.Bytes(), "", "\t")
		if err != nil {
			log.Panicf("marshalIndent error %s", err)
		}
		_, errStdout = io.Copy(stdout, bytes.NewReader(outBuf.Bytes()))
	}
	_, errStderr = io.Copy(stderr, stderrIn)
	err = cmd.Wait()
	if err != nil {
		log.Panicf(fmt.Sprintf("cmd.Run() failed with %s\n", err))
	}
	if errStdout != nil || errStderr != nil {
		log.Panicf("failed to capture stdout or stderr\n")
	}
}
