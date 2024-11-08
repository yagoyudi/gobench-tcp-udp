//go:build mage

package main

import "github.com/magefile/mage/sh"

var Default = Build

func Build() error {
	err := sh.RunV("go", "build", "-o", "bin/gobench", "./cmd/gobench")
	if err != nil {
		return err
	}
	return nil
}

func Test() error {
	err := sh.RunV("go", "test", "-v", "./...")
	if err != nil {
		return err
	}
	return nil
}

func Clean() error {
	err := sh.RunV("rm", "-rf", "bin")
	if err != nil {
		return err
	}
	return nil
}
