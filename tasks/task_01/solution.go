package main

import (
	"strings"
	"fmt"
)

func greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
