package command

import (
	"encoding/json"
	"fmt"
	"lockr/config"
	"os"
	"path/filepath"
)

func CreateCommand(envName string) {
	configPath := filepath.Join(LockrDir, ConfigFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to load config file: %v", err)
		os.Exit(1)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config file: %v", err)
		os.Exit(1)
	}
	for _, env := range config.Enviroments {
		if env == envName {
			fmt.Fprintf(os.Stderr, "Enviroment already exists \n: %v", envName)
			os.Exit(1)
		}
	}
	config.Enviroments = append(config.Enviroments, envName)

	updatedData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to marshal update config file: %v", err)
		os.Exit(1)

	}
	err = os.WriteFile(configPath, updatedData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write updated config file: %v", err)

	}

	folderPath := filepath.Join(LockrDir, "env", envName)
	CreateFile(folderPath)

	//localEnvFileName:= filepath.Join(lockrDir,"env",envName,".env")
	localEnvFileName := filepath.Join(LockrDir, "env", envName, ".env")
	err = os.WriteFile(localEnvFileName, []byte{}, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write env file: %v", err)

	}
	msg := fmt.Sprintf("Environment '%s' created successfully", envName)

	fmt.Print(msg)

}
