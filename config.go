package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure your sending and receiving directories",
	Run: func(cmd *cobra.Command, args []string) {
		// hande config command
	},
}

func ConfigDirChecker() (bool, error) {
	dataDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting config dir", err)
		return false, err
	}
	filePath := filepath.Join(dataDir, "qdm", "configdata.txt")
	_, err = os.Stat(filePath)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func getUsername() (string, error) {
	cmd := exec.Command("powershell.exe", "-Command", "Write-Host -NoNewLine $env:USERNAME")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting command output when determining Windows username.", err)
		return "", err
	}
	username := strings.TrimSpace(string(output))
	return username, nil
}

func winPathChecker() (bool, error) {
	dataDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting config dir", err)
		return false, err
	}
	configFilePath := filepath.Join(dataDir, "qdm", "configdata.txt")
	readData, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return false, err
	}
	_, err = os.Stat(string(readData))
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func CreateWinPath() {
	dataDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting config dir", err)
		os.Exit(1)
	}
	configFilePath := filepath.Join(dataDir, "qdm", "configdata.txt")
	readData, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	_ = exec.Command("mkdir", string(readData)).Run()
}

func CreateDefaultConfig() {
	dataDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting config dir", err)
		os.Exit(1)
	}
	configFilePath := filepath.Join(dataDir, "qdm", "configdata.txt")
	if err := os.MkdirAll(filepath.Dir(configFilePath), 0777); err != nil {
		fmt.Println("Error creating config dir", err)
		os.Exit(1)
	}
	username, err := getUsername()
	if err != nil {
		fmt.Println("Error getting Windows username", err)
		os.Exit(1)
	}
	fromWslPath := "/mnt/c/Users/" + username + "/Documents/fromWsl"
	fmt.Println(fromWslPath)
	_ = exec.Command("mkdir", fromWslPath).Run()
	writeDataPath := []byte(fromWslPath)
	if err := os.WriteFile(configFilePath, writeDataPath, 0777); err != nil {
		fmt.Println("Error creating config file:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created config file.")
}
