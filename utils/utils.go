package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"sigs.k8s.io/yaml"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type CustomError struct{}

func (m *CustomError) Error() string {
	return "Cannot open file, nil value returned."
}

func HandleStructured(jsonFile *os.File, fileType string) Users {
	byteValue, _ := io.ReadAll(jsonFile)

	var users Users

	if fileType == "json" {
		// unmarshall file  contents to users variable
		json.Unmarshal(byteValue, &users)
	} else if fileType == "yaml" {
		yaml.Unmarshal(byteValue, &users)
	} else {
		fmt.Println("Error: unsupported fileType")
	}

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + fmt.Sprint(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	return users
}

func HandleUnstructuredJson(jsonFile *os.File) []byte {
	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	j, err := json.Marshal(result)

	if err != nil {
		fmt.Println("Error: attempting to marshall json", err)
	}

	return j
}

func OpenJsonFileStructured(fileName string) (Users, error) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: opening json file: v%", fileName)
		return Users{}, err
	}

	if jsonFile != nil {
		fmt.Println("Opened file: v%", jsonFile.Name())

		defer jsonFile.Close()

		return HandleStructured(jsonFile, "json"), nil
	}

	return Users{}, &CustomError{}
}

func OpenJsonFileUnstructured(fileName string) ([]byte, error) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error: opening json file: %v", fileName)
		return nil, err
	}

	if jsonFile != nil {
		fmt.Printf("Opened file: %v", jsonFile.Name())

		defer jsonFile.Close()

		return HandleUnstructuredJson(jsonFile), nil
	}

	return nil, &CustomError{}
}

func OpenYamlFileStructured(fileName string) (Users, error) {
	yamlFile, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error: opening yaml file: %v", fileName)
		return Users{}, err
	}

	if yamlFile != nil {
		fmt.Printf("Opened file: %v", yamlFile.Name())

		defer yamlFile.Close()

		return HandleStructured(yamlFile, "yaml"), nil
	}

	return Users{}, &CustomError{}
}

func OpenYamlFileUnstructurd(fileName string) ([]byte, error) {
	return nil, &CustomError{}
}
