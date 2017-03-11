// reads url content and writes to a file, example: ./furl http://golang.com
package main

import (
  "fmt"
  "io"
  "bufio"
  "net/http"
  "os"
  "time"
  "strconv"
)
func main() {
  for _, url := range os.Args[1:]  {
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }

    // create file
    f, err := os.Create(strconv.FormatInt(time.Now().Unix(), 10) + "out.html")
    if err!=nil {
      fmt.Fprintf(os.Stderr, "fetch: creating file %v\n", err)
      os.Exit(1)
    }
    defer f.Close()

    _, err = io.Copy(bufio.NewWriter(f), bufio.NewReader(resp.Body))
    resp.Body.Close()
    if err != nil {
      fmt.Printf("fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
  }
}
