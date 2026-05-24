package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LoadRecipients(filepath string) error {
	p, err := os.Open(filepath)
	if err != nil {
		return err
	}

	data, err :=csv.NewReader(p).ReadAll()
	if err != nil {
		return err
	}

	for _, mails := range data[1:]{
		fmt.Println(mails)
	}
	return nil

}
