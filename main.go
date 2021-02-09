package main

import (
	"fmt"
	"io/ioutil"

	"./domain"
)

func main() {
	config := &domain.Config{}
	config.Data = make(map[string]interface{})

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", config)
	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}
	//config.Get("applications")
}
