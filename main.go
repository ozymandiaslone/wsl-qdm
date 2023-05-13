package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	configExists, err := ConfigDirChecker()
	if err != nil {
		log.Println(err)
		fmt.Println("This shit is broken, didn't work")
		os.Exit(1)
	}
	if !configExists {
		fmt.Println("No config directory found. Creating one...")
		CreateDefaultConfig()
	}
	winPathExists, err := winPathChecker()
	if err != nil {
		log.Println(err)
		fmt.Println("This shit is broken, didn't work")
		os.Exit(1)
	}
	if !winPathExists {
		fmt.Println("No windows path found. Creating one...")
		CreateWinPath()
	}
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
