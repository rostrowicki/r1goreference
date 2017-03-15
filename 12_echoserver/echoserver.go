// Usage: ./echoserver.exe 8000 echoserver.go
// will display content of file (echoserver.go) on localhost:8000
// you can use any URL form you need, e.g. localhost:8000/api/content - output is always file content
// might be helpful to mockup API calls without pinging prod server each time

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port int
var filename string

func main() {
	http.HandleFunc("/", handler)

	port = 8000
	if os.Args[1] != "" {
		port, _ = strconv.Atoi(os.Args[1])
	}

	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	filename = ""
	if os.Args[2] != "" {
		filename = os.Args[2]
	}

	fileContent := ""
	if filename != "" {
		fileContent = getFile(filename)
		fmt.Fprint(w, fileContent)
	} else {
		fmt.Fprintf(w, "URL.Path = %q/n", r.URL.Path)
	}

}

func getFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	f.Close()

	return string(buf.Bytes())
}
