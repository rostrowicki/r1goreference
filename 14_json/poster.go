// reading information about a movie from internet database
// usage: poster [Title Of The Movie]
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const SvcUrl = "http://www.omdbapi.com/" // http://www.omdbapi.com/?t=ghostbusters
const templ = `
---------------------------------------
Title:			{{.Title}}
Year:			{{.Year}}
Rated:			{{.Rated}}
Released:		{{.Released}}
Length (min):		{{.Runtime}}
Genre:			{{.Genre}}
Director:		{{.Director}}
Writer:			{{.Writer}}
Actors:			{{.Actors}}
Plot:			{{.Plot}}
Language:		{{.Language}}
Country:		{{.Country}}
Awards:			{{.Awards}}
Poster:			{{.Poster}}{{range .Ratings}}
Source/Rating:		{{.Source}} / {{.Value}}{{end}}
Metascore:		{{.Metascore}}
ImdbRating:		{{.ImdbRating}}
ImdbVotes:		{{.ImdbVotes}}
ImdbID:			{{.ImdbID}}
Dvd Release:		{{.Dvd}}
BoxOffice:		{{.BoxOffice}}
Production:		{{.Production}}
Website:		{{.Website}}
`

type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Dvd        string `json:"dvd_released_date"`
	BoxOffice  string
	Production string
	Website    string
}

type Rating struct {
	Source string
	Value  string
}

func main() {
	result, err := SearchMovie(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	report, err := template.New("report").Funcs(template.FuncMap{}).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

func SearchMovie(title []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(title, "+"))
	resp, err := http.Get(SvcUrl + "?t=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP error code: %s\n%s\n", resp.StatusCode, resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
