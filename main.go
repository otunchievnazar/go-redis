package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SomeData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	som := setSomeData("Flora", 25)
	fmt.Printf("DEBUG: som = %+v\n", som)

	jsonData, err := json.Marshal(som)
	if err != nil {
		http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
		return
	}
	fmt.Printf("DEBUG: jsonData = %s\n", jsonData)
	io.WriteString(w, string(jsonData))
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func setSomeData(newName string, newAge int) *SomeData {
	return &SomeData{
		Name: newName,
		Age:  newAge,
	}
}

func main() {
	som := setSomeData("Flora", 26)
	fmt.Println(som)
	fmt.Println(som.Name, som.Age)
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
