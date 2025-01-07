package utils

import "indexer/models"

func ProcessEmailDirectory(directory string, batchSize int, IndexEmailsToOpenObserve func([]models.Email) error) error {
	emailsChan, consumerDone := StartBatchConsumer(batchSize, IndexEmailsToOpenObserve)

	err := WalkDirectoryAndParse(directory, emailsChan)
	if err != nil {
		return err
	}

	close(emailsChan)
	consumerDone.Wait()

	return err
}
