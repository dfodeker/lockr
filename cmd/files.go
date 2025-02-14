package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
