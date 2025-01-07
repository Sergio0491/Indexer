package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"indexer/models"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	openObserveSearchURL string
	openObserveAuthUser  string
	openObserveAuthPass  string
	httpClient           *http.Client
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env. Usando variables de entorno del sistema.")
	}

	openObserveSearchURL = os.Getenv("OpenObserveSearchURL")
	openObserveAuthUser = os.Getenv("OpenObserveAuthUser")
	openObserveAuthPass = os.Getenv("OpenObserveAuthPass")

	if openObserveSearchURL == "" || openObserveAuthUser == "" || openObserveAuthPass == "" {
		log.Fatal("Faltan variables de entorno necesarias: OpenObserveSearchURL, OpenObserveAuthUser o OpenObserveAuthPass")
	}

	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func IndexEmailsToOpenObserve(emails []models.Email) error {
	emailsJSON, err := json.Marshal(emails)
	if err != nil {
		return fmt.Errorf("error serializando email: %v", err)
	}

	req, err := http.NewRequest("POST", openObserveSearchURL, bytes.NewReader(emailsJSON))
	if err != nil {
		return fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(openObserveAuthUser, openObserveAuthPass)

	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err = httpClient.Do(req)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("error enviando solicitud después de 3 intentos: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error leyendo respuesta: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error en OpenObserveSearch: %s, body: %s", resp.Status, string(body))
	}

	return nil
}
