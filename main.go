package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// .lockr folder name
const lockrDir string = ".lockr"

// Name of config file for Lockr
const configFile string = "config.json"

const gitIgnore string = ".gitignore"

type Config struct {
	Enviroments []string `json:"enviroments"`
	ActiveEnv   string   `json:"active_env"`
}

func main() {

	if len(os.Args) > 1 {
		subcmd := os.Args[1]
		switch subcmd {
		case "init":
			initCmd()
		default:
			err := fmt.Errorf("Unknown Command :%s", subcmd)
			fmt.Println(err.Error())
		}
	} else {
		msg := fmt.Sprintf("Welcome to Lockr! Use 'lockr init' to get started ")
		fmt.Println(msg)
	}

}

// Generates a .lockr file if it does not exist
// and add it as a line to a gitignore if present
// handle errors
func initCmd() (string, error) {
	cwd, err := os.Getwd()
	successMsg := fmt.Sprintf("Initialized Lockr 🚀 in %v", cwd)
	if err != nil {
		log.Fatal("Failed to get file path")
	}
	if _, err := os.Stat(lockrDir); err == nil {
		// for redabliity file = dir,

		//remove file and recreate it
		//using the os.MkdirAll will not let the user know if file exists
		//and will do nth if it does
		err := os.RemoveAll(lockrDir)
		if err != nil {
			log.Fatal("Failed to reinitialize lockr in ", cwd)
		}

		// createFile(cwd)
		// createConfig()
		fmt.Printf("Reinitialized existing Lockr vault in %v\n", cwd)
		reinitMsg := fmt.Sprintf("Reinitialized existing Lockr vault in %v\n", cwd)
		successMsg = reinitMsg
	}
	createFile(cwd)
	createConfig()
	appendGitIgnore()
	return successMsg, nil

}

// creates a default config file in the .lockr directory
func createConfig() error {
	configPath := filepath.Join(lockrDir, configFile)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		//should not overwrite file
		config := Config{
			Enviroments: []string{"default"},
			ActiveEnv:   "default",
		}
		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			//meaning error happs
			log.Fatal("Error Occured during config file creation")
			return err
		}
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			log.Fatal("Error Occured during config file write")
			return err
		}

	}
	return nil

}
func createFile(s string) (string, error) {
	err := os.MkdirAll(lockrDir, 0755)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	msg := fmt.Sprintf("File created in %v", s)
	return msg, nil
}

// appendGitIgnore
func appendGitIgnore() error {
	data, err := os.OpenFile(gitIgnore, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	file, err := os.Open(gitIgnore)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore for reading: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, lockrDir) {
			// fmt.Printf("found line %v", line)
			return nil
		}

	}
	data.WriteString("\n" + lockrDir + "\n")
	// fmt.Println("Successfully added to git ignore")
	return nil

}
