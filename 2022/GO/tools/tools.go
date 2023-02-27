package tools

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(filename string) (content []string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content

}

func Find(arr []string, str string) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}

	return false

}
