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
	Enviroments []string `json:"environments"`
	ActiveEnv   string   `json:"active_env"`
}

func main() {
	if len(os.Args) > 1 {
		// fullCmd := strings.Join(os.Args[1:], " ")
	Loop:
		for i := 1; i < len(os.Args); i++ {
			subCmd := os.Args[i]
			switch subCmd {
			case "init":
				initCmd()
			case "env":
				//flag check
				// if strings.Contains(fullCmd, "-c") {
				// 	fmt.Println("using the -c flag")
				// }

				if i+2 < len(os.Args) {
					switch os.Args[i+1] {
					case "create":
						envCmd(os.Args[i+2])
					case "switch":
						switchCmd(os.Args[i+2])
					default:
						fmt.Printf("Error: 'create' requires another argument")
					}

				} else {
					fmt.Printf("Error: 'create' requires an argument")
				}
				break Loop
			case "list":
				listCmd()
			default:
				err := fmt.Errorf("Unknown Command :%s", subCmd)
				fmt.Println(err.Error())
			}

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

	if err != nil {
		log.Fatal("Failed to get file path")
	}
	successMsg := fmt.Sprintf("Initialized Lockr 🚀 in %v", cwd)
	if _, err := os.Stat(lockrDir); err == nil {
		// for redabliity file = dir,

		//remove file and recreate it
		//using the os.MkdirAll will not let the user know if file exists
		//and will do nth if it does
		err := os.RemoveAll(lockrDir)
		if err != nil {
			log.Fatal("Failed to reinitialize lockr in ", cwd)
		}
		fmt.Printf("Reinitialized existing Lockr vault in %v\n", cwd)
		reinitMsg := fmt.Sprintf("Reinitialized existing Lockr vault in %v\n", cwd)
		successMsg = reinitMsg
	}
	createFile(lockrDir)
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
func createFile(folderName string) (string, error) {
	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	msg := fmt.Sprintf("Created %v vault directory", folderName)
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

// Creates a new environment (similar to creating a new branch in Git).
// (lockr env create <environment-name>)
func envCmd(envName string) (string, error) {

	configPath := filepath.Join(lockrDir, configFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("unable to load config file: %w", err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", fmt.Errorf("unable to parse config file: %w", err)
	}
	for _, env := range config.Enviroments {

		if env == envName {
			fmt.Printf("'%s' enviroment already exists \n", envName)
			return "", fmt.Errorf("enviroment '%s' already exists \n", envName)
		}
	}
	config.Enviroments = append(config.Enviroments, envName)

	updatedData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return "", fmt.Errorf("unable to marshal update config %w", err)
	}
	err = os.WriteFile(configPath, updatedData, 0644)
	if err != nil {
		return "", fmt.Errorf("unable to write updated config file: %w", err)
	}

	folderPath := filepath.Join(lockrDir, "env", envName)
	createFile(folderPath)

	//localEnvFileName:= filepath.Join(lockrDir,"env",envName,".env")
	localEnvFileName := filepath.Join(lockrDir, "env", envName, ".env")
	err = os.WriteFile(localEnvFileName, []byte{}, 0644)
	if err != nil {
		return "", fmt.Errorf("unable to write env file: %w", err)
	}
	return fmt.Sprintf("Environment '%s' created successfully", envName), nil
}

// ```bash
// lockr env create <environment-name>
// ```
// . Create and Switch to a New Environment
// This command creates a new environment and switches to it immediately (similar to git checkout -b).

// Switches to an already created enviroment
func switchCmd(envName string) (string, error) {
	//return early if there isnt a file in env
	//filePath.Join(lockrDir,"env",envName)
	envPath := filepath.Join(lockrDir, "env", envName)
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return "", fmt.Errorf("unable to locate enviroment: %w", err)
	}
	configPath := filepath.Join(lockrDir, configFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("unable to load config file: %w", err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", fmt.Errorf("Unable to parse config file: %w", err)
	}

	if config.ActiveEnv == envName {
		fmt.Printf("environment '%s' is already the active environment", envName)
		return "", fmt.Errorf("environment '%s' is already the active environment", envName)
	}
	config.ActiveEnv = envName
	updatedData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return "", fmt.Errorf("unable to marshal update config %w", err)
	}

	err = os.WriteFile(configPath, updatedData, 0644)
	if err != nil {
		return "", fmt.Errorf("unable to write updated config file: %w", err)
	}
	fmt.Printf("updated env\n")
	return "", nil
}

func listCmd() (string, error) {
	configPath := filepath.Join(lockrDir, configFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("unable to load config file: %w", err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", fmt.Errorf("Unable to parse config file: %w", err)
	}
	for _, env := range config.Enviroments {
		fmt.Printf("%v\n", env)
	}

	return "", nil
}
