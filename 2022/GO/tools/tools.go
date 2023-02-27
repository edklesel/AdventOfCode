package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("Line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
