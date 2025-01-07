package utils

import (
	"indexer/models"
	"log"
	"sync"
)

func StartBatchConsumer(batchSize int, indexFunc func([]models.Email) error) (chan models.Email, *sync.WaitGroup) {
	emailsChan := make(chan models.Email, batchSize*2)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		var batch []models.Email
		for email := range emailsChan {
			batch = append(batch, email)
			if len(batch) == batchSize {
				if err := indexFunc(batch); err != nil {
					log.Printf("Error enviando batch a OpenObserve: %v", err)
				}
				batch = nil
			}
		}
		if len(batch) > 0 {
			if err := indexFunc(batch); err != nil {
				log.Printf("Error enviando último batch a OpenObserve: %v", err)
			}
		}
	}()

	return emailsChan, &wg
}
