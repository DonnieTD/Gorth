package utils

import (
	"fmt"
	"os/exec"
)

func RunCMD(app string, arg1 string, arg2 string, arg3 string, arg4 string) {
	cmd := exec.Command(app, arg1, arg2, arg3, arg4)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
