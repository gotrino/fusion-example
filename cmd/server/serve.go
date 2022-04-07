package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()
	path = filepath.Join(path, "build")
	fmt.Println(path)
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir(path)))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
