package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Payload struct {
	Payload string
}

func main() {
	fmt.Println("testing")
	if len(os.Args) > 1 {
		argsWithoutProg := os.Args[1:]
		args := argsWithoutProg[0]
		p := Payload{}
		err := json.Unmarshal([]byte(args), &p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(args)
		fmt.Println(p)
	} else {
		fmt.Println("no args provided")
	}
}
