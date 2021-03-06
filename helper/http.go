package helper

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func InterfaceToRequest(data interface{}) (io.Reader, error) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func ResponseToInterface(response []byte, data interface{}) error {
	err := json.Unmarshal(response, &data)
	if err != nil {
		log.Fatal("ResponseToInterface:", err)
		return err
	}
	return nil
}

func Fetch(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("NewRequest:", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Fetch Do:", err)
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Fetch ReadAll:", err)
		return nil, err
	}

	return responseBody, nil
}
