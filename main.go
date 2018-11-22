package main

import (
	"net/http"
	"log"
	"encoding/json"
	"bytes"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	} 

	message := map[string]interface{}{
		"email": os.Args[1],
		"password":  os.Args[2],
		"undelete":  "false",
		"captcha_key":  "null",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://discordapp.com/api/v6/auth/login", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["token"] == nil {
		fmt.Println("Failed to get token.")
	} else {
		tokenik := fmt.Sprint(result["token"])
		fmt.Println(tokenik)

		file, err := os.Create("result.txt")
		if err != nil {
			log.Fatal("Cannot create file with token.", err)
		}
		file.WriteString(tokenik);
		defer file.Close()
	}
}
