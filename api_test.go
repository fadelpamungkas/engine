package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func sendGet() {
	// Get (GET http://localhost:8000/)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "http://localhost:8000/", nil)

	// Fetch Request
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Failure : ", err)
		fmt.Println("Failure : ", resp)
	}

	// Read Response Body
	// respBody, _ := io.ReadAll(resp.Body)

	// Display Results
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
}

func sendGetUsingCollectrows() {
	// Get using CollectRows (GET http://localhost:8000/row)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "http://localhost:8000/row", nil)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
		fmt.Println("Failure : ", resp)
	}

	// Read Response Body
	// respBody, _ := io.ReadAll(resp.Body)

	// Display Results
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
}

func sendInsertNewRow(name string, balance int, currency string) {
	// Insert new row (POST http://localhost:8000/insert)

	q := fmt.Sprintf(`{"name": "%s","balance": %d,"currency": "%s"}`, name, balance, currency)

	json := []byte(q)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "http://localhost:8000/insert", body)

	// Headers
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
		fmt.Println("Failure : ", resp)
	}

	// Read Response Body
	// respBody, _ := io.ReadAll(resp.Body)

	// Display Results
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
}

func TestInsertApi(t *testing.T) {
	testCases := []struct {
		name     string
		balance  int
		currency string
	}{
		{
			name:     "Pamungkas",
			balance:  1430000,
			currency: "IDR",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			sendInsertNewRow(tC.name, tC.balance, tC.currency)
		})
	}
}

func BenchmarkInsertApi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sendInsertNewRow("Bukan Fadel", 5400000, "USD")
	}
}

func BenchmarkGetRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sendGet()
	}
}

func BenchmarkGetCollectRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sendGetUsingCollectrows()
	}
}
