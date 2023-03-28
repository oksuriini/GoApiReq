package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type Response struct {
	Rand string `json:"rand"`
}

func main() {
	http.HandleFunc("/rand", retRand)
	http.ListenAndServe(":8080", nil)
}

func retRand(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	min := query.Get("min")
	max := query.Get("max")
	if max == "" {
		max = "100"
	}
	if min == "" {
		min = "1"
	}
	var result Response = Response{Rand: fmt.Sprint(getRand(max, min))}
	marshalled, _ := json.Marshal(result)
	fmt.Fprint(res, string(marshalled))
}

func getRand(max string, min string) (result int) {
	rmax, _ := strconv.Atoi(max)
	rmin, _ := strconv.Atoi(min)

	result = rand.Intn((rmax-rmin)+1) + rmin
	return
}
