package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/resend/resend-go/v2"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/send_email", SendEmail)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

type EmailRequest struct {
	SenderEmail string `json:"sender"`
	ReceiverEmail string `json:"receiver"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}

func SendEmail(w http.ResponseWriter, r *http.Request){
	var req EmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var apiKey = goDotEnvVariable("RESEND_API_KEY")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    req.SenderEmail,
		To:      []string{req.ReceiverEmail},
		Subject: req.Subject,
		Html:    req.Body,
	}

	_, err = client.Emails.Send(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func goDotEnvVariable(key string) string {
  return os.Getenv(key)
}