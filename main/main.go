package main

import (
	"fmt"
	"log"

	"example.com/transformer"
)

func printOutput(o string) {
		fmt.Println(o)
}

func getYAMLFromJSON() []byte {
	j := []byte(`{"name": "John", "age": 30}`)
	yaml, err := transformer.ToYaml(j)

	if err != nil {
		log.Fatal(err)
	}

	return yaml
}

func getJSONFromYAML() []byte {
	y := []byte(`{"name": "John", "age": 30}`)
	json, err := transformer.ToJson(y)

	if err != nil {
		log.Fatal(err)
	}

	return json
}

func main() {

	log.SetPrefix("Transformer: ")
	log.SetFlags(0)

	yaml := getYAMLFromJSON()
	printOutput(string(yaml))

	json := getJSONFromYAML()
	printOutput(string(json))
}
