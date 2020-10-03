package main

import (
	"fmt"
	"os/exec"
)

func main() {
	doTests := exec.Command("go", "test", "-coverprofile=coverage.out")
	if err := doTests.Run(); err != nil {
		fmt.Println("can not test")
	}

	showCoverage := exec.Command("go", "tool", "cover", "-html=coverage.out")
	if err := showCoverage.Run(); err != nil {
		fmt.Println("can not show coverage")
	}
}
