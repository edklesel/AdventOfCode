package tools

import (
	"bufio"
	"fmt"
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
		fmt.Printf("Line: %s\n", line)
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content

}
