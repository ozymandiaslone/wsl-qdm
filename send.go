package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func sendToWin(path string) {
	fmt.Println("send to win")
	dataDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Config error matey.")
		panic(err)
	}
	configFilePath := filepath.Join(dataDir, "qdm", "configdata.txt")
	readData, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading config file.")
		panic(err)
	}
	fmt.Println("Path: ", path)
	fmt.Println("ReadData: ", string(readData))
	if err := exec.Command("cp", "-r", path, string(readData)).Run(); err != nil {
		fmt.Println("Error running command to send file")
		panic(err)
	}
}
func sendCmd() *cobra.Command {
	sendCmd := &cobra.Command{
		Use:   "send <path>",
		Short: "Send a file to the configured directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("SEND COMMAND CALLED")
			path := args[0]
			fmt.Println("Path1: ", path)
			absPath, err := filepath.Abs(path)
			if err != nil {
				fmt.Println("path error matey.")
				panic(err)
			}
			fmt.Println("absPath: ", absPath)
			fmt.Println("Sending file(s)...")
			//cmdpath := filepath.Join(absPath, path)
			sendToWin(absPath)
		},
	}
	return sendCmd
}
