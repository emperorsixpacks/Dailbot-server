package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if pathStr, err := os.Getwd(); err == nil {
		fmt.Println(filepath.Dir(pathStr))
	}
}
