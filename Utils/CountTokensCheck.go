package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CountTokensCheck(current_tokens int, count int, inputPath string,fnName string) {
	if current_tokens != count {
		abs, err := filepath.Abs(inputPath)
		if err == nil {
			fmt.Printf("Error in: %v\nUpdate CURRENT_OPCOUNT %v() \n", abs,fnName)
		}
		os.Exit(1)
	}
}
