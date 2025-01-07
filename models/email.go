package models

type Email struct {
	MessageId               string `json:"message-id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mime-version"`
	ContentType             string `json:"content-type"`
	ContentTransferEncoding string `json:"content-transfer-encoding"`
	XFrom                   string `json:"x-from"`
	XTo                     string `json:"x-to"`
	XCc                     string `json:"x-cc"`
	XBcc                    string `json:"x-bcc"`
	XFolder                 string `json:"x-folder"`
	XOrigin                 string `json:"x-origin"`
	XFileName               string `json:"x-fileName"`
	Body                    string `json:"body"`
}
