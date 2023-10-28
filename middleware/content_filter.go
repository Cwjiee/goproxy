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


func ParseJson(filepath string) (*CensorList, error){
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

func FilterContent(body io.Reader) (io.Reader,int, error){
	bodyBytes, err := handlers.ReadResponse(body)
	if err != nil {
		return nil, 0, err	
	}

	data, err := ParseJson("config/filter_content.json")
	if err != nil {
		return nil, 0, err
	}

	var censorList []string = data.CensorList

	filteredContent := string(bodyBytes)
	for _, word := range censorList {
		filteredContent = strings.ReplaceAll(filteredContent, word, "****")
	}

	contentLength := len(filteredContent)

	return strings.NewReader(filteredContent), contentLength, nil
}
