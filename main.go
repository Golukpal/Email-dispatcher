package main

import (
	"fmt"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	recipientChan := make(chan Recipient)
	var wg sync.WaitGroup

	workerCount := 5
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go EmailWorker(i, recipientChan, &wg)
	}

	err := LoadRecipients("./mails.csv", recipientChan)
	if err != nil {
		fmt.Printf("Error loading recipients", err)
	}
	
	wg.Wait()
	fmt.Println("All emails sent successfully")
}
