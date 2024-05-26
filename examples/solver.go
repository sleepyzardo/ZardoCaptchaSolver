package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

func main() {
    apiUrl := "http://45.140.188.54:6200/solve"
    apiKey := "your api key"
    uid := "123456789"

    // Prepare the JSON payload
    payload := map[string]string{
        "apiKey": apiKey,
        "uid":    uid,
    }
    jsonData, err := json.Marshal(payload)
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        os.Exit(1)
    }

    // Create a new HTTP POST request
    req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error creating request:", err)
        os.Exit(1)
    }
    req.Header.Set("Content-Type", "application/json")

    // Execute the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    // Read and print the response
    if resp.StatusCode == http.StatusOK {
        var result map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&result)
        fmt.Println("Response:", result)
    } else {
        fmt.Printf("Error: received status code %d\n", resp.StatusCode)
    }
}
