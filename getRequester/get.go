package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8080/rand?max=100&min=50")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(data, data)
	fmt.Println(string(data))
}
