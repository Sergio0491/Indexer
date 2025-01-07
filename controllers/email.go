package controllers

import (
	"indexer/services"
	"indexer/utils"
	"log"
	"time"
)

func IndexEmails(directory string) error {
	startTime := time.Now()
	log.Println("Indexación iniciada.")

	err := utils.ProcessEmailDirectory(directory, 500, services.IndexEmailsToOpenObserve)

	if err != nil {
		return err
	}

	log.Println("Indexación terminada.")
	elapsedTime := time.Since(startTime)
	log.Printf("Tiempo total transcurrido: %s\n", elapsedTime)
	return nil
}
