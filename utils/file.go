package utils

import (
	"bufio"
	"log"
	"os"
)

func GetFileContents(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	contents := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		contents = append(contents, line)
	}
	return contents
}
