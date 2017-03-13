// image!
// minimal plus request counting
// try: http://localhost:8000/display?cycles=13
package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "sync"
  "image"
  "image/color"
  "image/gif"
  "io"
  "math"
  "math/rand"
  "strconv"
)

var mu sync.Mutex
var count int

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/count", counter)
  http.HandleFunc("/display", displayImage)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Number of requests: %d\n", count)
  mu.Unlock()
}

func displayImage(w http.ResponseWriter, r *http.Request) {

  m, _ := url.ParseQuery(r.URL.RawQuery)
  c, err := strconv.ParseFloat(m["cycles"][0], 64)

  if err != nil {
    panic(err)
  }

  if c<1 {
    c=5
  }

  mu.Lock()
  count++
  mu.Unlock()

  lissajous(w, c) // displaying anim gif in the browser
}

// generating anim gif and writing to destination
func lissajous(out io.Writer, cycles float64) {

  var palette = []color.Color{color.White, color.Black}

  const (
    res = 0.001
    size = 100
    nframes = 64
    delay = 8
    whiteIndex = 0
    blackIndex = 1
  )
  freq := rand.Float64() * 3.0
  anim := gif.GIF{LoopCount : nframes}
  phase := 0.1
  for i := 0; i < nframes; i++ {
    rect := image.Rect(0,0, 2*size+1, 2*size+1)
    img := image.NewPaletted( rect, palette)
    for t := 0.0; t<cycles*2*math.Pi; t += res {
      x := math.Sin(t)
      y := math.Sin(t * freq + phase)
      img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, delay)
    anim.Image = append(anim.Image, img)
  }
  gif.EncodeAll(out, &anim)
}
