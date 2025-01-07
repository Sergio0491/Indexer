package utils

import (
	"bufio"
	"fmt"
	"indexer/models"
	"io"
	"os"
	"strings"
)

func ParseEmailFile(path string) (models.Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return models.Email{}, err
	}
	defer file.Close()

	email := models.Email{}
	var body strings.Builder
	isBody := false

	scanner := bufio.NewScanner(file)

	const maxTokenSize = 1024 * 1024 // 1MB
	scanner.Buffer(make([]byte, 0, 64*1024), maxTokenSize)

	for scanner.Scan() {
		line := scanner.Text()
		if !isBody {
			if line == "" {
				isBody = true
				continue
			}
			processHeader(line, &email)
		} else {
			body.WriteString(line)
			body.WriteByte('\n')
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		return models.Email{}, fmt.Errorf("error escaneando archivo %s: %w", path, err)
	}

	email.Body = body.String()

	return email, nil
}

func processHeader(line string, email *models.Email) {
	headerMapping := map[string]func(string){
		"Message-ID: ":                func(value string) { email.MessageId = value },
		"Date: ":                      func(value string) { email.Date = value },
		"From: ":                      func(value string) { email.From = value },
		"Subject: ":                   func(value string) { email.Subject = value },
		"Mime-Version: ":              func(value string) { email.MimeVersion = value },
		"Content-Type: ":              func(value string) { email.ContentType = value },
		"Content-Transfer-Encoding: ": func(value string) { email.ContentTransferEncoding = value },
		"X-From: ":                    func(value string) { email.XFrom = value },
		"X-To: ":                      func(value string) { email.XTo = value },
		"X-cc: ":                      func(value string) { email.XCc = value },
		"X-bcc: ":                     func(value string) { email.XBcc = value },
		"X-Folder: ":                  func(value string) { email.XFolder = value },
		"X-Origin: ":                  func(value string) { email.XOrigin = value },
		"X-FileName: ":                func(value string) { email.XFileName = value },
	}

	for prefix, assignFunc := range headerMapping {
		if strings.HasPrefix(line, prefix) {
			trimmedValue := strings.TrimSpace(strings.TrimPrefix(line, prefix))
			assignFunc(trimmedValue)
			break
		}
	}
}
