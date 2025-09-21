package main

import (
	"bufio"
	"os"
	"strings"
)

var GREENLIGHT_DB_DSN string

// get data from .env
func init() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "GREENLIGHT_DB_DSN=") {
			GREENLIGHT_DB_DSN = strings.ReplaceAll(s, "GREENLIGHT_DB_DSN=", "")
		}
	}
}
