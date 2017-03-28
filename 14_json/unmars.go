package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Year      int `json:"Birth Date"`
}

var people = []Person{
	{FirstName: "Robert", LastName: "Sharp", Year: 1979},
	{FirstName: "John", LastName: "Doe", Year: 1987},
}

var guys = []Person{}

func main() {
	text := "!Hello world!"

	// create file
	f, err := os.Create("unmars.txt")
	if err != nil {
		log.Fatalf("Error while creating file...\n %s", err)
	}
	defer f.Close()

	// marshal ( struct -> json )
	data, err := json.Marshal(people)

	if err != nil {
		log.Fatalf("Error while marshalling..,\n %s", err)
	}
	text = fmt.Sprintf("%s\n", data)

	// write json to file
	w := bufio.NewWriter(f)
	w.WriteString(text)
	w.Flush()
	io.WriteString(os.Stdout, "Saved: "+text+"\n")

	// read string from file
	jsondata, err := ioutil.ReadFile(f.Name())
	text = fmt.Sprintf("%s", jsondata)

	// unmarshal ( struct <- json )
	fmt.Println("Unmarshalling...\n")
	if err := json.Unmarshal(jsondata, &guys); err != nil {
		log.Fatalf("Error while unmarshalling...\n %s", err)
	}
	fmt.Printf("%s\n", people)
	fmt.Printf("%s\n", guys)

	fmt.Println("Finished!")
}
