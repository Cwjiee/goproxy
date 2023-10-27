package middleware

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Cwjiee/goproxy/handlers"
)

type CensorList struct {
	CensorList []string `json:"words"`	
}


func ParseJson(filepath string, body io.Reader) (*CensorList, error){
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening json file")
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading json file")
	}
	
	var censorList CensorList

	err = json.Unmarshal(data, &censorList)
	if err != nil {
		log.Fatalf("Error unmarshalling json %v", err)
	}

	return &censorList, nil
}

func FilterContent(body io.Reader) (io.Reader, error){
	bodyBytes, err := handlers.ReadResponse(body)
	if err != nil {
		return nil, err	
	}

	data, err := ParseJson("config/filter_content.json", body)
	if err != nil {
		return nil, err
	}

	var censorList []string = data.CensorList

	for _, word := range censorList {
		bodyBytes = []byte(strings.ReplaceAll(string(bodyBytes), word, "****"))
	}

	return strings.NewReader(string(bodyBytes)), nil
}
