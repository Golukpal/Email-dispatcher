package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"text/template"
	"time"
)


func ExecuteTemplate(recipient Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, recipient)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}


func EmailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done() 

	smtpHost := "localhost"
	smtpPort := "1025" 

	for recipient := range ch {
		msgStr, err := ExecuteTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d: Error formatting email for %s\n", id, recipient.Email)
			continue
		}

		message := []byte(msgStr)

		
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "admin@mycompany.com", []string{recipient.Email}, message)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Worker %d: Sent email to %s\n", id, recipient.Email)
		
		
		time.Sleep(50 * time.Millisecond)
	}
}