package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/jtepe/powidl/pod"
)

func main() {
	episodes, err := pod.ParseFeed(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse feed: %v", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, e := range episodes {
		wg.Add(1)
		go func(e *pod.Episode) {
			defer wg.Done()
			fmt.Printf("downloading %q ...\n", e.Title)
			err := pod.Download(e)
			if err != nil {
				fmt.Fprintf(os.Stderr, "download of %q failed: %v\n", e.Title, err)
				return
			}
			fmt.Printf("finished %q with size %d\n", e.Title, e.Bytes)
		}(e)
	}

	wg.Wait()
}
