package main

import (
	"fmt"
	"os"

	"quantumwizard.hu/qwsg/internal/releasepublication"
)

const maxToolInput = 2 << 20

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "release-index:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 3 && len(args) != 4 {
		return fmt.Errorf("usage: qwsg-release-index generate INPUT OUTPUT | assemble INPUT SIGNATURE OUTPUT | verify INPUT CHECKPOINT")
	}
	input, err := readBounded(args[1])
	if err != nil {
		return err
	}
	var output []byte
	switch args[0] {
	case "generate":
		if len(args) != 3 {
			return fmt.Errorf("generate requires INPUT OUTPUT")
		}
		output, err = releasepublication.Generate(input)
	case "assemble":
		if len(args) != 4 {
			return fmt.Errorf("assemble requires INPUT SIGNATURE OUTPUT")
		}
		var signature []byte
		signature, err = readBounded(args[2])
		if err == nil {
			output, err = releasepublication.Assemble(input, signature, "qwsg-community-release-2026-01")
		}
	case "verify":
		if len(args) != 3 {
			return fmt.Errorf("verify requires INPUT CHECKPOINT")
		}
		output, err = releasepublication.BuildCheckpoint(input)
	default:
		return fmt.Errorf("unknown operation")
	}
	if err != nil {
		return err
	}
	output = append(output, '\n')
	return writeExclusive(args[len(args)-1], output)
}

func readBounded(name string) ([]byte, error) {
	info, err := os.Lstat(name)
	if err != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 || info.Size() > maxToolInput {
		return nil, fmt.Errorf("unsafe or oversized input")
	}
	return os.ReadFile(name)
}

func writeExclusive(name string, content []byte) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	if _, err = file.Write(content); err == nil {
		err = file.Sync()
	}
	closeErr := file.Close()
	if err != nil {
		return err
	}
	return closeErr
}
