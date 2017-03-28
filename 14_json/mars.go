package main

import (
    "fmt"
    "encoding/json"
    "log"
)

type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json:"color,omitempty"`
    Actors[] string
}

type Book struct {
    Title string
    Year int
    Authors[] string
}

var movies = [] Movie {
    { Title: "MovieA", Year: 1942, Color: false, Actors: [] string { "John", "Alice" } }, { Title: "MovieB", Year: 1999, Color: true, Actors: [] string { "Alan", "Laura" } },
}

var books = [] Book {
    { Title: "DDD", Year: 2016, Authors: [] string { "V.Vernon" } }, { Title: "Programming Go", Year: 2015, Authors: [] string { "Donovan", "Kernigan" } },
}


func main() {
    fmt.Println(movies)
    fmt.Println(books)

    fmt.Println("--- --- ---")

    // JSON formatting
    data, err: = json.Marshal(movies)
    if (err != nil) {
        log.Fatalf("Marshal failed: %s", err)
    }
    fmt.Printf("%s\n", data)
}