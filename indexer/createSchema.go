package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("loading JSON Schema Object")
	jsonFileName := "schema.json"
	urlZinchSearch := "http://localhost:4080/api/index"
	bytesFile, err := readJSONFile(jsonFileName)
	if err != nil {
		panic(err)
	}
	_ = sendData(urlZinchSearch, bytesFile)
}

func readJSONFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func sendData(url string, jsonData []byte) error {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	username := "admin"
	password := "Complexpass#123"
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle the response as needed
	//fmt.Println(resp)
	return nil
}
