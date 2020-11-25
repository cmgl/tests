package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Got environment variable:", os.Getenv("GITHUB_TOKEN")[0:20])
}
