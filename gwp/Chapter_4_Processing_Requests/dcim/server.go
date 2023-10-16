package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const token = "ssss"

func main() {

	url := "https://dcim-ito.hit.edu.cn/api/dcim/devices/"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	count := data["count"].(float64)

	fmt.Println(count)

}
