package utils

import (
	"fmt"
	"os/exec"
)

func RunCMD(cmd string) {
	fmt.Println(cmd)
	_, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        fmt.Println("Error running command: "+cmd)
    }
}
