package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-getter"
)

func main() {
	// Targeted Alpine release
	releaseTarget := "https://dl-cdn.alpinelinux.org/alpine/v3.15/releases/aarch64/alpine-rpi-3.15.0-aarch64.tar.gz"

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Print("Failed to detect current directory\n")
		os.Exit(1)
	}
	fileDest := currentDir + "/alpine-rpi"

	fmt.Printf("Downloading release into %v\n", fileDest)
	getter.Get(fileDest, releaseTarget)
	fmt.Print("Done.")
}
