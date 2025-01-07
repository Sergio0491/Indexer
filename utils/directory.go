package utils

import (
	"indexer/models"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func WalkDirectoryAndParse(directory string, emailsChan chan<- models.Email) error {
	maxGoroutines := runtime.NumCPU()
	sem := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	processFile := func(path string) {
		defer wg.Done()
		defer func() { <-sem }()

		email, err := ParseEmailFile(path)
		if err != nil {
			log.Printf("Error parseando archivo %s: %v", path, err)
			return
		}
		emailsChan <- email
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		sem <- struct{}{}
		wg.Add(1)
		go processFile(path)

		return nil
	})

	wg.Wait()

	return err
}
