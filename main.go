package main

import (
	"fmt"
	"log"
	"os"

	"example.com/transformer"
	"example.com/utils"
)

func printOutput(o string) {
	fmt.Println(o)
}

func getYAMLFromJSON(j []byte) []byte {
	// j := []byte(`{"name": "John", "age": 30}`)
	yaml, err := transformer.ToYaml(j)

	if err != nil {
		log.Fatal(err)
	}

	return yaml
}

func getJSONFromYAML(y []byte) []byte {
	json, err := transformer.ToJson(y)

	if err != nil {
		log.Fatal(err)
	}

	return json
}

func main() {

	log.SetPrefix("Transformer: ")
	log.SetFlags(0)

	file := "data/users.json"

	jsonData, err := utils.OpenJsonFileUnstructured(file)

	if err != nil {
		log.Println("Error: reading contents of: ", file)
		log.Fatal(err)
		os.Exit(1)
	}

	yaml := getYAMLFromJSON(jsonData)
	printOutput(string(yaml))

	// json := getJSONFromYAML()
	// printOutput(string(json))
}
