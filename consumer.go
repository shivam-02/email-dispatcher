package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		msg, err := excecuteTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker:%d Error parsing template for %s", id, recipient.Email)
			continue
		}
		fmt.Printf("Worker %d: Sending email to %s\n", id, recipient.Email)
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "john@doe.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d: Sent email to %s\n", id, recipient.Email)

	}
}
