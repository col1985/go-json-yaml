package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func HandleStructuredJson(jsonFile *os.File) Users {
	byteValue, _ := io.ReadAll(jsonFile)

	var users Users

	// unmarshall file  contents to users variable
	json.Unmarshal(byteValue, &users)

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

		return HandleStructuredJson(jsonFile), nil
	}

	return Users{}, &CustomError{}
}

func OpenJsonFileUnstructured(fileName string) ([]byte, error) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: opening json file: v%", fileName)
		return nil, err
	}

	if jsonFile != nil {
		fmt.Println("Opened file: v%", jsonFile.Name())

		defer jsonFile.Close()

		return HandleUnstructuredJson(jsonFile), nil
	}

	return nil, &CustomError{}
}
