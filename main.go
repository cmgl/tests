package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Got environment variable:", os.Getenv("TEST_VARIABLE")[0:20])
}
