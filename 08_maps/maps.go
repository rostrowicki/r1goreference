package main

import "fmt"

func main() {
  fmt.Printf("Maps...\n")


  var countryCapitalMap map[string]string
  /* creating map */
  countryCapitalMap = make(map[string]string)

  countryCapitalMap["France"] = "Paris"
  countryCapitalMap["Italy"] = "Rome"
  countryCapitalMap["Japan"] = "Tokyo"
  countryCapitalMap["India"] = "New Delhi"

  /* display */
  // for country := range countryCapitalMap {
  //   fmt.Println("Capital of", country, "is", countryCapitalMap[country])
  // }
  printMap(countryCapitalMap)

  /* delete */
  delete(countryCapitalMap, "France")
  fmt.Println("Entry deleted...")

  countryCapitalMap["Poland"] = "Warsaw"
  fmt.Println("Capital added...")

  printMap(countryCapitalMap)

}

func printMap(countryCapitalMap map[string]string) {
  /* display */
  for country := range countryCapitalMap {
    fmt.Println("Capital of", country, "is", countryCapitalMap[country])
  }
}
