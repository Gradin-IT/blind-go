package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Hunter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

var hunters = []Hunter{}

func loadHunters() {
	response, err := http.Get("https://raw.githubusercontent.com/GradinIT/hunters/develop/hunters.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(responseData), &hunters)
	for _, hunter := range hunters {
		//log.Println(hunter.Name)
		hunters = append(hunters, hunter)
	}
}
func GetHunters() []Hunter {
	loadHunters()
	return hunters
}
