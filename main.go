package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/mitchellh/go-ps"
)

func main() {
	for pr, err := ps.FindProcess(os.Getppid()); err == nil && pr != nil; pr, err = ps.FindProcess(pr.PPid()) {
		switch pr.Executable() {
		case "zed":
			run("zed", "--wait")
			return
		case "Code Helper": // VSC
			run("code", "--wait")
			return
		}
	}
	// fallback
	run(os.Getenv("EDITOR"))
}

func run(name string, args ...string) {
	cmd := exec.Command(name, append(args, os.Args[1:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s: %s", cmd.String(), err)
	}
}
