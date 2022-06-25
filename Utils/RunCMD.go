package utils

import (
	"fmt"
	"os/exec"
)

func RunCMD(cmd string) {
	_, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        fmt.Println("some error found")
    }
}
