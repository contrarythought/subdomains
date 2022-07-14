package main

import (
	"fmt"
	"os"
	"sync"
	"tricoder/json_structs"
)

const (
	BASE_URL = "https://crt.sh"
)

// TODO
func enumerate_targets(targets []string) json_structs.CRT {
	var cert json_structs.CRT
	var wg sync.WaitGroup

	for _, target := range targets {
		wg.Add(1)
		go func(target string) {
			// connect to API
			// unmarshal data into struct
			wg.Done()
		}(target)
	}

	// wait for workers to finish
	wg.Wait()

	return cert
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter targets to scan")
		os.Exit(1)
	}

	//certs := enumerate_targets(os.Args[1:])
}
