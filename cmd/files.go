package command

import (
	"bufio"
	"encoding/json"
	"fmt"
	config "lockr/config"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateFile(folderName string) (string, error) {
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
	data, err := os.OpenFile(GitIgnore, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	file, err := os.Open(GitIgnore)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore for reading: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, LockrDir) {
			// fmt.Printf("found line %v", line)
			return nil
		}

	}
	data.WriteString("\n" + LockrDir + "\n")
	// fmt.Println("Successfully added to git ignore")
	return nil

}

func createConfig() error {
	configPath := filepath.Join(LockrDir, ConfigFile)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		//should not overwrite file
		config := config.Config{
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
