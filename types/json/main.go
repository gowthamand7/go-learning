package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func main() {

	resp := Response{
		Status: "OK",
		Data:   "Hello World",
	}

	var resp2 Response

	j, _ := json.Marshal(resp)
	fmt.Printf("json %s \n", j)

	json.Unmarshal(j, &resp2)

	fmt.Printf("struct %+v \n", resp2)

	resp3 := &Response{
		Status: "OK",
		Data:   "Hello World",
	}

	fmt.Printf("pointer %#v \n", resp3)

	fmt.Printf("address %v", *(resp3))

}
